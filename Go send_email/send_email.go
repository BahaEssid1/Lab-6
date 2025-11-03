// package main

// import (
// 	"fmt"
// 	"net/smtp"
// )

// func main() {
// 	// === Configuration ===
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"
// 	sender := "essidbaha24@gmail.com"
// 	password := "ywkjpcnvakvbpnhr"
// 	receiver := "essidbaha18@gmail.com"

// 	subject := "Lab 6 ALERT: Test Email"
// 	message := "This is a test alert message from your Go script to baha essid lab 6."

// 	body := "Subject: " + subject + "\r\n\r\n" + message

// 	auth := smtp.PlainAuth("", sender, password, smtpHost)

// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, []string{receiver}, []byte(body))
// 	if err != nil {
// 		fmt.Println("[ERROR] Failed to send email:", err)
// 		return
// 	}

// 	fmt.Println("[OK] Email sent successfully!")
// }
