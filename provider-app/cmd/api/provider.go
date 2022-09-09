package api

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"net"
)

type GreeterProvider struct {
	GreeterServer
}

func getIpv4() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func (s *GreeterProvider) SayHello(ctx context.Context, in *HelloRequest) (*User, error) {
	logger.Infof("Dubbo3 GreeterProvider get user name = %s\n", in.Name)
	return &User{Name: "Hello " + in.Name, Id: getIpv4(), Age: 21}, nil
}
