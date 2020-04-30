package main

import (
	"fmt"
	"net"
	_ "net/http/pprof"
	"os"

	consul "github.com/hashicorp/consul/api"
	"github.com/prometheus/common/log"
)

func RegisterService(id, name, address string, port int) error {
	reg := &consul.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Port:    port,
		Address: address,
		// HTTP Service
		Check: &consul.AgentServiceCheck{
			Interval: "5s",
			HTTP:     fmt.Sprintf("http://%s:%d/", address, port),
		},
	}
	config := consul.DefaultConfig()
	config.Address = "https://<A>.<dns>.com:8500"
	config.Token = "<token with service_prefix with write permission use "" for allowing all services to access>"
	client, err := consul.NewClient(config)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceRegister(reg)
	return err
}

func main() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Infoln("Begin to regist service on consul")
				if ipnet.IP.String() != "127.0.0.1" {
					regErr := RegisterService(ipnet.IP.String(), "nodes", ipnet.IP.String(), 9100)
					if regErr != nil {
						fmt.Println(regErr)
					}
					log.Infoln("Regist complete!")
				}
			}
		}
	}
}
