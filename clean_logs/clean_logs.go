package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Exact path to your logs folder
	folder := `C:\Users\Baha\Desktop\Developpement full stack\Lab 6\logs`

	now := time.Now()

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		// Only files with .log extension
		if !info.IsDir() && filepath.Ext(path) == ".log" {
			age := now.Sub(info.ModTime()).Hours() / 24 // Age in days
			if age > 7 {
				fmt.Println("[DELETE]", path)
				os.Remove(path)
			} else {
				fmt.Println("[KEEP]", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}
