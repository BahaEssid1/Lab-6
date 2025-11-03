package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://www.netacad.com/"

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[ERROR] Cannot reach %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("[OK] %s is reachable!\n", url)
	} else {
		fmt.Printf("[ERROR] %s returned status code %d\n", url, resp.StatusCode)
	}
}
