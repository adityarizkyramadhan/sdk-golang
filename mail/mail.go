package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

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
