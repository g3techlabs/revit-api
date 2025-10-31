package mail

import (
	"fmt"
	"html/template"

	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/infra/mail/templates"
	"github.com/wneessen/go-mail"
)

// TODO: Integrar fila com Redis para envio de emails

type IEmailService interface {
	SendPassResetEmail(to, name, deepLink string, expiration int) error
}

type EmailService struct{}

func NewEmailService() IEmailService {
	return &EmailService{}
}

var appName = config.Get("APP_NAME")
var mailUser = config.Get("MAIL_USER")
var mailPassword = config.Get("MAIL_PASS")
var mailPort = config.GetIntVariable("MAIL_PORT")
var mailHost = config.Get("MAIL_HOST")

func (es *EmailService) buildClient() (*mail.Client, error) {
	tlsPolicy := es.chooseTlsPolicy()

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

func (es *EmailService) chooseTlsPolicy() mail.TLSPolicy {
	switch mailPort {
	case 465:
		return mail.TLSMandatory
	case 587:
		return mail.TLSMandatory
	default:
		return mail.TLSOpportunistic
	}
}

func (es *EmailService) SendPassResetEmail(destinatary, name, deepLink string, expiration int) error {

	template, err := template.ParseFS(templates.FS, "reset_password.tmpl")
	if err != nil {
		return err
	}

	msg := mail.NewMsg()
	if err := msg.FromFormat(appName+" App", mailUser); err != nil {
		return err
	}
	if err := msg.To(destinatary); err != nil {
		return err
	}

	msg.Subject("Redefinição de senha")

	data := map[string]any{
		"Name":       name,
		"DeepLink":   deepLink,
		"AppName":    appName,
		"Expiration": expiration,
	}
	if err := msg.SetBodyHTMLTemplate(template, data); err != nil {
		return err
	}

	client, err := es.buildClient()
	if err != nil {
		return err
	}

	if err := client.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
