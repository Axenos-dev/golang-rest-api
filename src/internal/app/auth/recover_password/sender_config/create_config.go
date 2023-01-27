package sender_config

type SMTPConfig struct {
	sender_email    string
	sender_password string
}

func (config *SMTPConfig) CreateConfig(email string, password string) {
	config.sender_email = email
	config.sender_password = password
}
