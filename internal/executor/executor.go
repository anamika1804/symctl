package executor

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/SymmetricalAI/symctl/internal/logger"
)

func Execute(plugin string, args []string) {
	pluginExecutable := fmt.Sprintf("symctl-%s", plugin)

	logger.Debugf("Plugin executable: %s\n", pluginExecutable)
	logger.Debugf("Plugin arguments: %v\n", args)

	cmd := exec.Command(pluginExecutable, args...)
	logger.Debugf("Executing command: %v\n", cmd)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating StdoutPipe for Cmd: %v\n", err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error creating StderrPipe for Cmd: %v\n", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting Cmd: %v\n", err)
	}

	multiReader := io.MultiReader(stdoutPipe, stderrPipe)
	scanner := bufio.NewScanner(multiReader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalf("Cmd finished with error: %v\n", err)
	}
}
