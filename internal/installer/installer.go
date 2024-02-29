package installer

import "fmt"

func Install(address string, args []string) {
	fmt.Printf("Installing plugin from %s with args %v\n", address, args)
}
