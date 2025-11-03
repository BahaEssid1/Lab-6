package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Folder to backup
	folder := `C:\Users\Baha\Desktop\Developpement full stack\Lab 6\logs`
	// Output ZIP file
	zipFile := `C:\Users\Baha\Desktop\Developpement full stack\Lab 6\backupp.zip`

	zipf, err := os.Create(zipFile)
	if err != nil {
		fmt.Println("Error creating ZIP:", err)
		return
	}
	defer zipf.Close()

	archive := zip.NewWriter(zipf)
	defer archive.Close()

	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking path:", err)
			return err
		}
		if info.IsDir() {
			return nil // skip directories
		}

		// Relative path inside the zip
		relPath, err := filepath.Rel(folder, path)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		w, err := archive.Create(relPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(w, file)
		return err
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Backup created at", zipFile)
}
