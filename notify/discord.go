package notify

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/bwmarrin/discordgo"
	jsoniter "github.com/json-iterator/go"
)


var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Discord struct {
	ID        string
	Token     string
	ChannelID string
	PostName  string
	Message   string
}

func (d Discord) Endpoint() string {
	return discordgo.EndpointWebhookToken(d.ID, d.Token)
}

func (d Discord) Send() error {
	if err := d.check();err != nil {
		return err
	}

	var params discordgo.WebhookParams
	params.Username = d.PostName
	params.Content = d.MaxChars()
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", d.Endpoint(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// 投稿未遂チェック
	if res.StatusCode != 200 && res.StatusCode != 204 {
		return errors.New(res.Status)
	}

	return nil
}

func (p Discord) check() error {
	if p.ID == ""  {
		return errors.New("has not id")
	} else if  p.Token == "" {
		return errors.New("has not token")
	}
	return nil
}

func (p Discord) MaxChars() string {
	// the message contents (up to 2000 characters)
	if 2000 < len([]rune(p.Message))  {
		return string([]rune(p.Message)[:2000])
	}

	return p.Message
}
