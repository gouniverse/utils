package utils

import (
	"log"
	"net/smtp"
	"os"

	"github.com/darkoatanasovski/htmltags"
	"github.com/jordan-wright/email"
)

// EmailSend sends an email
func EmailSend(from string, to []string, subject string, htmlMessage string) (bool, error) {
	//drvr := os.Getenv("MAIL_DRIVER")
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	user := os.Getenv("MAIL_USERNAME")
	pass := os.Getenv("MAIL_PASSWORD")
	addr := host + ":" + port

	nodes, errStripped := htmltags.Strip(htmlMessage, []string{}, true)

	textMessage := ""

	if errStripped == nil {
		//nodes.Elements   //HTML nodes structure of type *html.Node
		textMessage = nodes.ToString() //returns stripped HTML string
	}

	e := email.NewEmail()
	e.From = from
	e.To = to
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = subject
	e.Text = []byte(textMessage)
	e.HTML = []byte(htmlMessage)
	err := e.Send(addr, smtp.PlainAuth("", user, pass, host))

	// Choose auth method and set it up
	// auth := smtp.PlainAuth("", user, pass, host)

	// mime := "MIME-version: 1.0;"
	// mime += "\r\n"
	// mime += "Content-Type: text/html; charset=\"UTF-8\";"

	// msg := "To: " + to[0]
	// msg += "\r\n"
	// msg += "Subject: " + subject
	// msg += "\r\n"
	// msg += mime
	// msg += "\r\n"
	// msg += "\r\n"
	// msg += message
	// msg += "\r\n"

	// err := smtp.SendMail(addr, auth, from, to, []byte(msg))
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}
