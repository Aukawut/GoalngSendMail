package mail

import (
	"crypto/tls"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

func SendMail(c *fiber.Ctx) error {
	// Email server configuration
	server := "10.145.0.250"
	port := 25 // Port 25 for insecure, or change to 587 for TLS
	sender := "it-system@prospira.local"
	pwd := "Psth@min135"

	// Recipient email and message content
	to := "akawut.kamesuwan@prospira.com"
	subject := "Test Email from Go"
	body := "This is a test email sent using Go."

	// Create a new email message
	mail := gomail.NewMessage()
	mail.SetHeader("From", sender)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", body)

	// Configure the SMTP dialer
	dialer := gomail.NewDialer(server, port, sender, pwd)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Equivalent to `rejectUnauthorized: false`

	// Send the email
	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
		return c.JSON(fiber.Map{"err": true, "msg": err.Error()})
	}
	return c.JSON(fiber.Map{"err": false, "msg": "Email sent successfully!"})

}

func SendEMailToApprover(c *fiber.Ctx) error {
	// Email server configuration
	server := "10.145.0.250"
	port := 25 // Port 25 for insecure, or change to 587 for TLS
	sender := "it-system@prospira.local"
	pwd := "Psth@min135"

	// Recipient email and message content
	to := "akawut.kamesuwan@prospira.com"
	subject := "Test Email from Go"
	body := "<p>This is a test email sent using Go.</p>"

	// Create a new email message
	mail := gomail.NewMessage()
	mail.SetHeader("From", sender)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	// Configure the SMTP dialer
	dialer := gomail.NewDialer(server, port, sender, pwd)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Equivalent to `rejectUnauthorized: false`

	// Send the email
	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
		return c.JSON(fiber.Map{"err": true, "msg": err.Error()})
	}
	return c.JSON(fiber.Map{"err": false, "msg": "Email sent successfully!"})

}
