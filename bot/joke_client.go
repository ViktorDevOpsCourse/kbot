package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
