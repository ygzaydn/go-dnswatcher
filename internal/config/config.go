package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DNSEntity struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Config struct {
	PollingInterval int         `yaml:"polling_interval"`
	DnsServers      []DNSEntity `yaml:"dns_servers"`
}

func LoadConfig(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func ParseConfig(cfg *Config) string {
	result := fmt.Sprintf("Polling Interval: %d seconds\n", cfg.PollingInterval)
	result += "DNS Servers:\n"
	for i, server := range cfg.DnsServers {
		result += fmt.Sprintf("  %d. %s:%d\n", i+1, server.IP, server.Port)
	}
	return result
}
