package bot

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"kbot/config"
	"log"
	"net/http"
	"time"
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

	b.Use(middleware.Logger())

	return &TelegramBot{
		bot: b,
	}
}

func (t *TelegramBot) InitHandlers() {
	fmt.Printf("kbot %s started", config.Version)
	t.bot.Handle("/start", func(m telebot.Context) error {

		menu := t.bot.NewMarkup()
		// Reply buttons.
		btnJoke := menu.Text("Tell me a joke")
		menu.Reply(menu.Row(btnJoke))

		return m.Send("Welcome to fanny Kbot!", menu)
	})

	t.bot.Handle(telebot.OnText, func(m telebot.Context) error {
		switch m.Text() {
		case "Tell me a joke":
			joke, err := GetJoke()
			if err != nil {
				log.Println(err)
				return m.Send(fmt.Sprintf("Sorry some error occured when we try to tell you joke." +
					"Not worry we fixing it! ;)"))
			}
			return m.Send(joke)
		default:
			return m.Send("Sorry but i can only tell you joke")
		}
	})

	t.bot.Start()
}

type JokeResponse struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Joke     string `json:"joke"`
	Delivery string `json:"delivery"`
	Flags    struct {
		NSFW      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
	ID   int    `json:"id"`
	Safe bool   `json:"safe"`
	Lang string `json:"lang"`
}

func GetJoke() (string, error) {
	url := "https://v2.jokeapi.dev/joke/Any"

	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make the request: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to retrieve a joke. Status code: %d", response.StatusCode)
	}

	var jokeData JokeResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&jokeData)
	if err != nil {
		return "", fmt.Errorf("failed to decode the response: %s", err)
	}

	if jokeData.Error {
		return "", fmt.Errorf("failed get new joke from %s", url)
	}

	if jokeData.Type == "single" {
		return jokeData.Joke, nil
	}

	return fmt.Sprintf("%s\n%s", jokeData.Setup, jokeData.Delivery), nil
}
