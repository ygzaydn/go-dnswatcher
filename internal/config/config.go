package config

import (
	"os"

	"gopkg.in/yaml.v3"
)


type DNSEntity struct {
	IP string `yaml:"ip"`
	Port int `yaml:"port"`
}

type Config struct {
	PollingInterval int 	`yaml:"polling_interval"`
	DnsServers []DNSEntity	`yaml:"dns_servers"`
}

func LoadConfig(fileName string) (*Config,error){
	data,err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	
	cfg := &Config{}
	err = yaml.Unmarshal(data,cfg)

	if err != nil {
		return nil, err
	}

	return cfg,nil
}