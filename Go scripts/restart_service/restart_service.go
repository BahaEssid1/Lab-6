package main

import (
	"fmt"
	"os/exec"
)

func main() {
	service := "CloudflareWARP" // Exact service name

	cmd := exec.Command("powershell", "Restart-Service", service)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("[ERROR] Failed to restart %s: %v\n", service, err)
		return
	}

	fmt.Printf("[OK] %s restarted successfully.\n", service)
}
