package kpi

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/google/gopacket/layers"
)

type DNSMetrics struct {
	TotalQueries    uint64
	SuccessQueries  uint64
	FailedQueries   uint64
	QueryTypes      map[uint16]*atomic.Uint64
	ResponseLatency atomic.Int64
}

var GlobalMetrics = NewDNSMetrics()

func NewDNSMetrics() *DNSMetrics {
	m := &DNSMetrics{
		QueryTypes: make(map[uint16]*atomic.Uint64),
	}
	commonTypes := []uint16{1, 2, 5, 15, 28} // A, NS, CNAME, MX, AAAA
	for _, t := range commonTypes {
		m.QueryTypes[t] = &atomic.Uint64{}
	}
	return m
}

func GetMetrics() *DNSMetrics {
	return GlobalMetrics
}

func (m *DNSMetrics) IncrementQuery(queryType uint16) {
	atomic.AddUint64(&m.TotalQueries, 1)
	if counter, exists := m.QueryTypes[queryType]; exists {
		counter.Add(1)
	}
}

func (m *DNSMetrics) RecordResponse(success bool) {
	if success {
		atomic.AddUint64(&m.SuccessQueries, 1)
	} else {
		atomic.AddUint64(&m.FailedQueries, 1)
	}
}

func (m *DNSMetrics) PrintStats() {
	total := atomic.LoadUint64(&m.TotalQueries)
	success := atomic.LoadUint64(&m.SuccessQueries)
	failed := atomic.LoadUint64(&m.FailedQueries)

	log.Printf("=========================")
	log.Printf("=== DNS Statistics ===")
	log.Printf("Total Queries: %d", total)
	log.Printf("Successful Responses: %d", success)
	log.Printf("Failed Responses: %d", failed)
	log.Printf("Query Types:")
	for qType, counter := range m.QueryTypes {
		count := counter.Load()
		if count > 0 {
			log.Printf("  %s: %d", layers.DNSType(qType), count)
		}
	}
	log.Println("=========================")
}

func StatsString(m *DNSMetrics) string {
	total := atomic.LoadUint64(&m.TotalQueries)
	success := atomic.LoadUint64(&m.SuccessQueries)
	failed := atomic.LoadUint64(&m.FailedQueries)

	stats := "=== DNS Statistics ===\n"
	stats += fmt.Sprintf("Total Queries: %d\n", total)
	stats += fmt.Sprintf("Successful Responses: %d\n", success)
	stats += fmt.Sprintf("Failed Responses: %d\n", failed)
	stats += "Query Types:\n"
	for qType, counter := range m.QueryTypes {
		count := counter.Load()
		if count > 0 {
			stats += fmt.Sprintf("  %s: %d\n", layers.DNSType(qType), count)
		}
	}

	return stats
}
