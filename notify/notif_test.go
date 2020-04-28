package notify_test

import (
	"testing"
	"os"
	
	"github.com/go-numb/go-utilitys/notify"
)

func TestDiscord(t *testing.T) {
	discord := &notify.Discord{
		ID: os.Getenv("DISCORD_ID"),
		Token: os.Getenv("DISCORD_TOKEN"),
		ChannelID: "notif_bots",
		PostName: "test_bot",
		Message:"test",
	}
	t.Logf("%+v",discord)


	if err := discord.Send();err != nil {
		t.Fatal(err)
	}
}

func TestLine(t *testing.T) {
	line := &notify.Line{
		Token: os.Getenv("LINE_TOKEN"),
		Message:"test",
	}
	t.Logf("%+v",line)


	if err := line.Send();err != nil {
		t.Fatal(err)
	}
}


func TestBoth(t *testing.T) {
	both := make(map[string]notify.Notifyer)

	both["discord"] = &notify.Discord{
		ID: os.Getenv("DISCORD_ID"),
		Token: os.Getenv("DISCORD_TOKEN"),
		ChannelID: "notif_bots",
		PostName: "test_bot",
		Message:"test",
	}
	both["line"] = &notify.Line{
		Token: os.Getenv("LINE_TOKEN"),
		Message:"test",
	}


	for k := range both {
		if err := both[k].Send();err != nil {
			t.Fatal(err)
		}
	}
}