package cmd

import (
	"github.com/ct-core-standard/cmd/service"
	"github.com/ct-core-standard/pkg/log"
	"github.com/spf13/cobra"
	"os"
)

var logger = log.GetLogger()

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "application",
		Long:  "application",
		Run:  func(_ *cobra.Command, args []string) {},
	}
	rootCmd.AddCommand(service.Service)

	err := rootCmd.Execute()
	if err != nil {
		logger.Errorf("rootCmd: %v", err)
		os.Exit(1)
	}
	return err
}

