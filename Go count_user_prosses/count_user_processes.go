package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	username := "Baha" // Change to the user you want to check

	// Run 'tasklist /V'
	cmd := exec.Command("tasklist", "/V")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("[ERROR] Failed to get processes:", err)
	}

	lines := strings.Split(string(output), "\n")
	count := 0
	for _, line := range lines {
		if strings.Contains(line, username) {
			count++
		}
	}

	fmt.Printf("[OK] User '%s' has %d running processes.\n", username, count)
}
