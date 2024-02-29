package main

import (
	"fmt"
	"os"
)

func main() {
	println("Hello, World!")

	args := os.Args

	fmt.Println("Whole command-line: ", args)

	fmt.Println("Command-line arguments: ", args[1:])

	pluginExecutable := fmt.Sprintf("symctl-%s", os.Args[1])

	fmt.Println("Plugin executable: ", pluginExecutable)

	pluginArgs := os.Args[2:]

	fmt.Println("Plugin arguments: ", pluginArgs)
}
