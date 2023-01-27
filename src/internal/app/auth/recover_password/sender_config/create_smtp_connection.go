package sender_config

import "net/smtp"

func CreateConnection(config SMTPConfig) smtp.Auth {
	return smtp.PlainAuth("", config.sender_email, config.sender_password, "smtp.gmail.com")
}
