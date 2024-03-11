package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SymmetricalAI/symctl/internal/executor"
	"github.com/SymmetricalAI/symctl/internal/installer"
	"github.com/SymmetricalAI/symctl/internal/logger"
)

var (
	version = "xxx"
)

func init() {
	if os.Getenv("DEBUG") != "" {
		logger.Debug = true
	}
}

func main() {
	args := os.Args

	logger.Debugf("Whole command-line: %v\n", args)
	logger.Debugf("Command-line arguments: %v\n", args[1:])

	if len(args) < 2 {
		log.Fatalln("No command-line arguments provided.")
	}
	switch os.Args[1] {
	case "version":
		fmt.Printf("Version %s\n", version)
	case "install":
		if len(args) < 3 {
			log.Fatalln("No plugin address provided.")
			return
		}
		logger.Debugf("Installer called.\n")
		installer.Install(os.Args[2], os.Args[3:])
	default:
		logger.Debugf("Executor called.\n")
		executor.Execute(os.Args[1], os.Args[2:])
	}
}
