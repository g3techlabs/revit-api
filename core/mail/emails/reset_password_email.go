package emails

import (
	_ "embed"
	"html/template"

	"github.com/g3techlabs/revit-api/config"
	mailer "github.com/g3techlabs/revit-api/core/mail"
	"github.com/g3techlabs/revit-api/core/mail/templates"
	"github.com/wneessen/go-mail"
)

var appName = config.Get("APP_NAME")
var mailUser = config.Get("MAIL_USER")

func SendResetPasswordEmailService(destinatary, name, deepLink string, expiration int) error {

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

	client, err := mailer.BuildClient()
	if err != nil {
		return err
	}

	if err := client.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
