package mail

import (
	"fmt"
	"net/smtp"

	"github.com/cyan-store/hook/config"
	"github.com/cyan-store/hook/log"
)

func SendMail(status, msg string) error {
	auth := smtp.PlainAuth("", config.Data.Mail.User, config.Data.Mail.Password, config.Data.Mail.Host)

	if err := smtp.SendMail(
		fmt.Sprintf("%s:%d", config.Data.Mail.Host, config.Data.Mail.Port),
		auth, config.Data.Mail.User, []string{config.Data.ReportMail.Address},
		[]byte(fmt.Sprintf(
			"From: %s\r\n"+
				"To: %s\r\n"+
				"Subject: Webhook status: %s\r\n\r\n"+
				"Webhook report: %s\r\n",

			config.Data.Mail.User,
			config.Data.ReportMail.Address,
			status, msg,
		))); err != nil {

		log.Error.Println("[SendMail] Could not send email -", err)
		return err
	}

	return nil
}

func ReportSuccess(msg string) {
	if !config.Data.ReportMail.Success {
		return
	}

	go SendMail("SUCCESS", msg)
}

func ReportError(id, msg string, err error) {
	if !config.Data.ReportMail.Error {
		return
	}

	go SendMail("ERROR", fmt.Sprintf("%s\n\n%s - %s", id, msg, err))
}
