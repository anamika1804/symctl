package installer

import (
	"fmt"
	"os"
)

func Install(address string, args []string) {
	fmt.Printf("Installing plugin from %s with args %v\n", address, args)

	executablePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %s\n", err)
		return
	}
	fmt.Printf("Executable path: %s\n", executablePath)
}
