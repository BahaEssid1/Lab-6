package main

import (
	"fmt"
	"os/exec"
)

func main() {
	scriptPath := `C:\Users\Baha\Desktop\Developpement full stack\Lab 6\clean_logs.py`
	taskName := "Lab6DailyTask"

	// Command to create the scheduled task
	cmd := exec.Command("schtasks", "/Create",
		"/SC", "DAILY",
		"/TN", taskName,
		"/TR", fmt.Sprintf("python \"%s\"", scriptPath),
		"/ST", "02:00",
		"/F")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[ERROR] Failed to create scheduled task:", string(output))
		return
	}

	fmt.Println("[OK] Scheduled task created successfully!")
	fmt.Println(string(output))
}
