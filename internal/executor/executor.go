package executor

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func Execute(plugin string, args []string) {
	pluginExecutable := fmt.Sprintf("symctl-%s", plugin)

	fmt.Println("Plugin executable: ", pluginExecutable)

	fmt.Println("Plugin arguments: ", args)

	cmd := exec.Command(pluginExecutable, args...)
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
