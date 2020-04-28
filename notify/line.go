package notify

import (
	"fmt"
	"strings"
	"net/url"
	"errors"
	"net/http"
)



type Line struct {
	Token     string
	Message   string
	File string
}

func (p Line) Endpoint() string {
	return "https://notify-api.line.me/api/notify"
}


func (l Line) Send() error {
	if err := l.check();err != nil {
		return err
	}

	v := url.Values{}
	v.Set("message", l.MaxChars())
	if l.File != "" {
		v.Set("imageFile", "@"+l.File)
	}
	req, err := http.NewRequest("POST", l.Endpoint(), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s",l.Token))
	

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	return nil
}

func (p Line) check() error {
	if p.Token == "" {
		return errors.New("has not token")
	}
	return nil
}

func (p Line) MaxChars() string {
	// the message contents (up to 1000 characters)
	if 1000 < len([]rune(p.Message))  {
		return string([]rune(p.Message)[:1000])
	}

	return p.Message
}
