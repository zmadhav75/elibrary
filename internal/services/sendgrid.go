// internal/services/sendgrid.go
package services

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendAvailabilityEmail(userEmail, bookTitle string) error {
	from := mail.NewEmail("elibrary", "noreply@elibrary.com")
	to := mail.NewEmail("User", userEmail)
	subject := "Book Available Notification"
	plainTextContent := "Book %s is now available!"
	htmlContent := "<strong>Book %s is now available!</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	return err
}
