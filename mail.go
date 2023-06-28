package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"crypto/tls"
)

func main() {
	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpUsername := "your-email@gmail.com"
	smtpPassword := "your-password"

	// Sender and recipient email addresses
	from := "your-email@gmail.com"
	to := "recipient-email@example.com"

	// Email content
	subject := "Hello from Go"
	body := "This is the email body."

	// Create a new email message
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// Create the SMTP dialer
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// Send the email
	err := dialer.DialAndSend(message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
