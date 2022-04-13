package service

import (
	"context"
	server "github.com/ct-core-standard/pkg/protobuf"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

// RunGRPCGateway starts the invoke in process http gateway.
func RunGRPCGateway(ctx context.Context, addr string, srv server.AdListingServiceServer, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)
	server.RegisterAdListingServiceHandlerServer(ctx, mux, srv)
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		logger.Infof("Shutting down the http gateway server")
		if err := s.Shutdown(context.Background()); err != nil {
			logger.Errorf("Failed to shutdown http gateway server: %v", err)
		}
	}()
	logger.Infof("http gateway is running at port  %v", s.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logger.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}
