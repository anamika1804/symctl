package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
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

	cmd := exec.Command(pluginExecutable, pluginArgs...)
	fmt.Println("Executing command: ", cmd)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd:", err)
		return
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating StderrPipe for Cmd:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting Cmd:", err)
		return
	}

	multiReader := io.MultiReader(stdoutPipe, stderrPipe)
	scanner := bufio.NewScanner(multiReader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Cmd finished with error:", err)
		return
	}
}
