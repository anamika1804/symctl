package main

import (
	"fmt"
	"github.com/SymmetricalAI/symctl/internal/executor"
	"github.com/SymmetricalAI/symctl/internal/installer"
	"os"
)

func main() {
	println("Hello, World!")

	args := os.Args

	fmt.Println("Whole command-line: ", args)

	fmt.Println("Command-line arguments: ", args[1:])

	if len(args) < 2 {
		fmt.Println("No command-line arguments provided")
		return
	}
	switch os.Args[1] {
	case "version":
		fmt.Println("Version: TODO") //TODO implement
	case "install":
		if len(args) < 3 {
			fmt.Println("No plugin address provided")
			return
		}
		fmt.Println("Installer called")
		installer.Install(os.Args[2], os.Args[3:])
	default:
		fmt.Println("Executor called")
		executor.Execute(os.Args[1], os.Args[2:])
	}
}
