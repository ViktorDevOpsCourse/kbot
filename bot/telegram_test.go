package bot

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	tBot := New(os.Getenv("TELE_TOKEN"))

	if tBot.bot == nil {
		t.Error("Failed create new telegram bot")
	}
}
