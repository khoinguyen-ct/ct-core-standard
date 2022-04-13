package main

import (
	"github.com/ct-core-standard/cmd"
	"github.com/ct-core-standard/pkg/log"
)

var logger = log.GetLogger()

func main() {
	err := cmd.Execute()
	if err != nil {
		logger.Errorf("application execute err: %s", err)
	}
}
