package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

// const (
// 	CONFIG_SMTP_HOST     = "smtp.gmail.com"
// 	CONFIG_SMTP_PORT     = 587
// 	CONFIG_SENDER_NAME   = "Aditya Rizky Ramadhan <emailer2311@gmail.com>"
// 	CONFIG_AUTH_EMAIL    = "emailer2311@gmail.com"
// 	CONFIG_AUTH_PASSWORD = "jokgjhacdiqxjboo"
// )

// func main() {
// 	to := []string{"adityarizky2040@gmail.com"}
// 	cc := []string{"adityarizky1020@gmail.com"}
// 	subject := "Test mail"
// 	message := "WKOKOWKOKWOK BILANG SAYA KLO MASUK"

// 	err := sendMail(to, cc, subject, message)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	log.Println("Mail sent!")
// }

type mailer struct {
	SmtpHost   string
	SmtpPort   int
	SenderName string
	Email      string
	Password   string
}

func NewMailerOne(senderName, email, password, smtpHost string, smtpPort int) *mailer {
	return &mailer{
		SmtpHost:   smtpHost,
		SmtpPort:   smtpPort,
		SenderName: senderName,
		Email:      email,
		Password:   password,
	}
}

func (m *mailer) SendMail(to, cc []string, subject, message string) error {
	body := "From: " + m.SenderName + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", m.Email, m.Password, m.SmtpHost)
	smtpAddr := fmt.Sprintf("%s:%d", m.SmtpHost, m.SmtpPort)

	err := smtp.SendMail(smtpAddr, auth, m.Email, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
