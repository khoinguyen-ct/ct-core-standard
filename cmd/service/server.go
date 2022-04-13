package service

import (
	"context"
	"fmt"
	"github.com/ct-core-standard/app"
	"github.com/ct-core-standard/pkg/log"
	"github.com/spf13/cobra"
	"os"
)

var logger = log.GetLogger()

var Service = &cobra.Command{
	Use:   "service",
	Short: "API Command of service",
	Long:  "API Command of service",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof("Starting service...")
		RunServer()
	},
}

func RunServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := runServers(ctx)

	select {
	case err := <-errCh:
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runServers(ctx context.Context) <-chan error {
	ch := make(chan error, 3)
	srv := app.NewAdListingServiceServer(ctx)
	go func() {
		if err := RunGRPCServer(ctx, "tcp", ":8989", srv); err != nil {
			ch <- fmt.Errorf("cannot run grpc service: %v", err)
		}
	}()
	go func() {
		if err := RunGRPCGateway(ctx, ":8080", srv); err != nil {
			ch <- fmt.Errorf("cannot run in process gateway service: %v", err)
		}
	}()
	return ch
}
