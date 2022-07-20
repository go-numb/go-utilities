package notify_test

import (
	"os"
	"testing"
)

func TestDiscord(t *testing.T) {
	discord := &Discord{
		ID:        os.Getenv("DISCORD_ID"),
		Token:     os.Getenv("DISCORD_TOKEN"),
		ChannelID: "notif_bots",
		PostName:  "test_bot",
		Message:   "test",
	}
	t.Logf("%+v", discord)

	if err := discord.Send(); err != nil {
		t.Fatal(err)
	}
}

func TestLine(t *testing.T) {
	line := &Line{
		Token:   os.Getenv("LINE_TOKEN"),
		Message: "test",
	}
	t.Logf("%+v", line)

	if err := line.Send(); err != nil {
		t.Fatal(err)
	}
}

func TestBoth(t *testing.T) {
	both := make(map[string]Notifyer)

	both["discord"] = &Discord{
		ID:        os.Getenv("DISCORD_ID"),
		Token:     os.Getenv("DISCORD_TOKEN"),
		ChannelID: "notif_bots",
		PostName:  "test_bot",
		Message:   "test",
	}
	both["line"] = &Line{
		Token:   os.Getenv("LINE_TOKEN"),
		Message: "test",
	}

	for k := range both {
		if err := both[k].Send(); err != nil {
			t.Fatal(err)
		}
	}
}
