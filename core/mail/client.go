package mail

import (
	"fmt"

	"github.com/g3techlabs/revit-api/config"
	"github.com/wneessen/go-mail"
)

var mailUser = config.Get("MAIL_USER")
var mailPassword = config.Get("MAIL_PASS")
var mailPort = config.GetIntVariable("MAIL_PORT")
var mailHost = config.Get("MAIL_HOST")

func BuildClient() (*mail.Client, error) {
	tlsPolicy := chooseTlsPolicy()

	client, err := mail.NewClient(mailHost,
		mail.WithPort(mailPort),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(mailUser),
		mail.WithPassword(mailPassword),
		mail.WithTLSPolicy(tlsPolicy),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating SMTP client: %w", err)
	}

	return client, nil
}

func chooseTlsPolicy() mail.TLSPolicy {
	switch mailPort {
	case 465:
		return mail.TLSMandatory
	case 587:
		return mail.TLSMandatory
	default:
		return mail.TLSOpportunistic
	}
}
