package dnsHandler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"

	"github.com/ygzaydn/go-dnswatcher/internal/config"
)


func startDNSListener(server config.DNSEntity, wg* sync.WaitGroup){
	defer wg.Done()

	handle, err := pcap.OpenLive("any", 1600, true, pcap.BlockForever)
    if err != nil {
        log.Printf("Error opening device: %v", err)
        return
    }
    defer handle.Close()

	filter := fmt.Sprintf("udp port %v and (host %s)", server.Port, server.IP)

    if err := handle.SetBPFFilter(filter); err != nil {
        log.Printf("Error setting BPF filter: %v", err)
        return
    }

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    log.Printf("Starting DNS monitor on interface: any")

	for packet := range packetSource.Packets() {
        dnsLayer := packet.Layer(layers.LayerTypeDNS)
        if dnsLayer == nil {
            continue
        }

        dns, _ := dnsLayer.(*layers.DNS)
        srcIP := packet.NetworkLayer().NetworkFlow().Src()
        dstIP := packet.NetworkLayer().NetworkFlow().Dst()

        if dns.QR {
            log.Printf("DNS Response from %v to %v [%s]", srcIP, dstIP, dns.ResponseCode)
            for _, answer := range dns.Answers {
                log.Printf("  Answer: %s (%s) -> %v", 
                    string(answer.Name),
                    layers.DNSType(answer.Type),
                    answer.String(),
                )
            }
        } else {
            for _, question := range dns.Questions {
                log.Printf("DNS Query from %v to %v: %s (%s)", 
                    srcIP, dstIP,
                    string(question.Name),
                    layers.DNSType(question.Type),
                )
            }
        }
    }

}

func Start(cfg config.Config) {
	var wg sync.WaitGroup

	for _,server := range cfg.DnsServers {
		wg.Add(1)
		go startDNSListener(server, &wg)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		wg.Wait()
		cancel()
	}()

	<-ctx.Done()
	log.Println("Shutting down DNS listeners...")
	time.Sleep(time.Second)
}