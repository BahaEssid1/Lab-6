package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("[ERROR] Cannot get current user:", err)
		return
	}

	fmt.Println("[OK] Currently logged-in user:", u.Username)
}
