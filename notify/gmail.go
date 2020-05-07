package notify

import "net/smtp"

type Gmail struct {
	From     string
	Username string
	// Password is password/appword
	Password string
	To       string
	Subject  string
	Message  string
}

func (g Gmail) Endpoint() string {
	return "smtp.gmail.com"
}

func (g Gmail) Send() error {
	auth := smtp.PlainAuth("", g.Username, g.Password, g.Endpoint())
	if err := smtp.SendMail(g.Endpoint()+":587", auth, g.From, []string{g.To}, g.body()); err != nil {
		return err
	}
	return nil

	return nil
}

func (g Gmail) body() []byte {
	return []byte("To: " + g.To + "\r\n" +
		"Subject: " + g.Subject + "\r\n\r\n" +
		g.MaxChars() + "\r\n")
}

func (g Gmail) MaxChars() string {
	return g.Message
}
