package smtp

import (
	"fmt"
	"strconv"

	"github.com/ronaldyonatan/paramedics/pkg/utils"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	From       string
	To         []string
	Cc         []string
	Subject    string
	Body       string
	Attachment []string
}

func SendMail(req Mail) (err error) {
	mailer := gomail.NewMessage()
	mailer.SetHeader(
		"From",
		fmt.Sprintf("%s <%s>", req.From, utils.GetEnv("CONFIG_AUTH_EMAIL")),
	)
	for _, v := range req.To {
		mailer.SetHeader("To", v)
	}
	for _, v := range req.Cc {
		mailer.SetAddressHeader("Cc", v, v)
	}
	mailer.SetHeader("Subject", req.Subject)
	mailer.SetBody("text/html", req.Body)
	for _, v := range req.Attachment {
		mailer.Attach(v)
	}
	port, _ := strconv.Atoi(utils.GetEnv("CONFIG_SMTP_PORT"))
	dialer := gomail.NewDialer(
		utils.GetEnv("CONFIG_SMTP_HOST"),
		port,
		utils.GetEnv("CONFIG_AUTH_EMAIL"),
		utils.GetEnv("CONFIG_AUTH_PASSWORD"),
	)

	err = dialer.DialAndSend(mailer)

	if err != nil {
		return
	}
	return
}
