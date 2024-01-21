package bot

import (
	"context"
	"fmt"
	"gopkg.in/telebot.v3"
	"kbot/config"
	"log"
	"time"
)

const (
	JokeCommand = "Tell me joke"
)

type TelegramBot struct {
	bot *telebot.Bot
}

func New(token string) *TelegramBot {
	b, err := telebot.NewBot(telebot.Settings{
		URL:    "",
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
		return nil
	}

	//b.Use(middleware.Logger())

	return &TelegramBot{
		bot: b,
	}
}

func (t *TelegramBot) InitHandlers() {
	log.Printf("kbot %s started", config.Version)
	t.handleStart()
	t.handleReplayOnMessage()
}

func (t *TelegramBot) handleStart() {
	t.bot.Handle("/start", func(m telebot.Context) error {
		menu := t.bot.NewMarkup()
		btnJoke := menu.Text(JokeCommand)
		menu.ResizeKeyboard = true
		menu.Reply(menu.Row(btnJoke))
		return m.Send("Welcome to funny Kbot!", menu)
	})
}

func (t *TelegramBot) handleReplayOnMessage() {
	t.bot.Handle(telebot.OnText, func(m telebot.Context) error {
		switch m.Text() {
		case JokeCommand:
			defer pmetrics(context.TODO())
			joke, err := GetJoke()
			if err != nil {
				log.Println(err)
				return m.Send(fmt.Sprintf("Sorry some error occured when we try to tell you joke." +
					"Don't worry we are fixing it! ;)"))
			}

			log.Printf("New joke: '%s'", joke)

			return m.Send(joke)
		default:
			log.Printf("Uknow command to bot: '%s'", m.Text())
			return m.Send("Sorry but i can only tell you joke")
		}
	})
}

func (t *TelegramBot) StartBot() {
	t.bot.Start()
}
