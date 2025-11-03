package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type FileInfo struct {
	Name string
	Size int64
}

func main() {
	folder := `C:\Users\Baha\Desktop\Developpement full stack\Lab 6` // Change if needed

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	var fileInfos []FileInfo
	for _, f := range files {
		if !f.IsDir() {
			fileInfos = append(fileInfos, FileInfo{Name: f.Name(), Size: f.Size()})
		}
	}

	// Sort by size descending
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Size > fileInfos[j].Size
	})

	fmt.Println("[OK] 5 plus gros fichiers dans le dossier:")
	for i, f := range fileInfos {
		if i >= 5 {
			break
		}
		fmt.Printf("%s - %.2f KB\n", f.Name, float64(f.Size)/1024)
	}
}
