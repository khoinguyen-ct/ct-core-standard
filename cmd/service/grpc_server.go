package service

import (
	"context"
	server "github.com/ct-core-standard/pkg/protobuf"
	"google.golang.org/grpc"
	"net"
)

// RunGRPCServer starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func RunGRPCServer(ctx context.Context, network, address string, srv server.AdListingServiceServer) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			logger.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	server.RegisterAdListingServiceServer(s, srv)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	logger.Infof("grpc service is running at port %v", address)
	return s.Serve(l)
}
