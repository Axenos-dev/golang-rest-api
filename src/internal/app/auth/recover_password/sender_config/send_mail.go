package sender_config

import "net/smtp"

func Send_Mail(auth smtp.Auth, config SMTPConfig, target []string, message []byte) error {
	return smtp.SendMail("smtp.gmail.com:587", auth, config.sender_email, target, message)
}
