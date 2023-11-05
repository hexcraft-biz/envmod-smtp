package smtp

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"github.com/hexcraft-biz/misc/resp"
)

// ================================================================
//
// ================================================================
type Smtp struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DisplayName  string
	DisplayEmail string
}

func New() (*Smtp, *resp.Resp) {
	return &Smtp{
		Host:         os.Getenv("SMTP_HOST"),
		Port:         os.Getenv("SMTP_PORT"),
		Username:     os.Getenv("SMTP_USERNAME"),
		Password:     os.Getenv("SMTP_PASSWORD"),
		DisplayName:  os.Getenv("SMTP_DISPLAY_NAME"),
		DisplayEmail: os.Getenv("SMTP_DISPLAY_EMAIL"),
	}, nil
}

// ================================================================
//
// ================================================================
func (e Smtp) SendHTMLEmail(to []string, subject, body string) *resp.Resp {
	server := e.Host + ":" + e.Port
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)

	msgs := []string{
		`MIME-version: 1.0;`,
		`Content-Type: text/html; charset="UTF-8";`,
		fmt.Sprintf("From: %s <%s>", e.DisplayName, e.DisplayEmail),
		fmt.Sprintf("Subject: %s", subject),
		body,
	}

	err := smtp.SendMail(server, auth, e.Username, to, []byte(strings.Join(msgs, "\n")))
	return resp.NewError(http.StatusInternalServerError, err, nil)
}
