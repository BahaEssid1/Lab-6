// package main

// import (
// 	"fmt"

// 	"github.com/shirou/gopsutil/v3/disk"
// )

// func main() {
// 	path := `C:\` // Use "/" on Linux

// 	usage, err := disk.Usage(path)
// 	if err != nil {
// 		fmt.Println("[ERROR] Cannot get disk stats:", err)
// 		return
// 	}

// 	percentFree := (float64(usage.Free) / float64(usage.Total)) * 100
// 	fmt.Printf("Disk free: %.2f%%\n", percentFree)

// 	if percentFree < 10 {
// 		fmt.Println("[ALERT] Disk space is below 10%!")
// 	} else {
// 		fmt.Println("[OK] Disk space is sufficient.")
// 	}
// }
