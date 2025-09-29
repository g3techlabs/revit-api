package mail

import (
	"fmt"

	"github.com/g3techlabs/revit-api/config"
	"github.com/wneessen/go-mail"
)

var mailUser = config.Get("MAIL_USER")
var mailPassword = config.Get("MAIL_PASSWORD")
var mailPort = config.GetIntVariable("MAIL_PORT")
var mailHost = config.Get("MAIL_HOST")

func BuildClient() (*mail.Client, error) {
	client, err := mail.NewClient(mailHost, mail.WithPort(mailPort), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(mailUser), mail.WithPassword(mailPassword))
	if err != nil {
		return nil, fmt.Errorf("error creating SMTP client: %w", err)
	}

	return client, err
}
