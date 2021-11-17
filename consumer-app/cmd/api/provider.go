package api

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
)


type GreeterProvider struct {
	GreeterProviderBase
}

func (s *GreeterProvider) SayHello(ctx context.Context, in *HelloRequest) (*User, error) {
	logger.Infof("Dubbo3 GreeterProvider get user name = %s\n", in.Name)
	return &User{Name: "Hello " + in.Name, Id: "12345", Age: 21}, nil
}
