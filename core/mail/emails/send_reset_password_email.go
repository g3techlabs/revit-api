package emails

import (
	"fmt"
	"html/template"

	"github.com/g3techlabs/revit-api/config"
	mailer "github.com/g3techlabs/revit-api/core/mail"
	"github.com/wneessen/go-mail"
)

var appName = config.Get("APP_NAME")
var mailUser = config.Get("MAIL_USER")

func SendResetPasswordEmailService(destinatary, name, code string, expiration int) error {

	template, err := template.ParseFiles("templates/reset-password.tmpl")
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

	subject := fmt.Sprintf("Seu token para redefinir a senha é %s", code)
	msg.Subject(subject)

	data := map[string]any{
		"Name":       name,
		"Code":       code,
		"Expiration": expiration,
	}
	if err := msg.SetBodyHTMLTemplate(template, data); err != nil {
		return err
	}

	client, err := mailer.BuildClient()
	if err != nil {
		return err
	}

	if err := client.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
