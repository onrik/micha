package micha

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"
)

const (
	API_URL = "https://api.telegram.org/bot%s/%s"
)

type ApiResponse struct {
	Ok          bool            `json:"ok"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
	Result      json.RawMessage `json:"result"`
}

type Bot struct {
	token   string
	Me      User
	Timeout time.Duration
	Updates chan Update
}

// Create new bot instance
func NewBot(token string) (*Bot, error) {
	bot := Bot{
		token:   token,
		Timeout: 25 * time.Second,
		Updates: make(chan Update),
	}

	if me, err := bot.GetMe(); err != nil {
		return nil, err
	} else {
		bot.Me = *me
		return &bot, nil
	}
}

// Build url for API method
func (bot *Bot) buildUrl(method string) string {
	return fmt.Sprintf(API_URL, bot.token, method)
}

// Decode response body to ApiResponse object
func (bot *Bot) decodeResponse(data []byte) (*ApiResponse, error) {
	apiResponse := &ApiResponse{}
	if err := json.Unmarshal(data, apiResponse); err != nil {
		return nil, fmt.Errorf("Decode response error (%s)", err.Error())
	}

	if !apiResponse.Ok {
		return nil, fmt.Errorf("%d: %s", apiResponse.ErrorCode, apiResponse.Description)
	} else {
		return apiResponse, nil
	}
}

// Make GET request to Telegram API
func (bot *Bot) get(method string, params url.Values) (*ApiResponse, error) {
	response, err := get(bot.buildUrl(method) + "?" + params.Encode())
	if err != nil {
		return nil, err
	} else {
		return bot.decodeResponse(response)
	}
}

// Make POST request to Telegram API
func (bot *Bot) post(method string, data interface{}) (*ApiResponse, error) {
	response, err := post(bot.buildUrl(method), data)
	if err != nil {
		return nil, err
	} else {
		return bot.decodeResponse(response)
	}
}

// Use this method to receive incoming updates using long polling.
// An Array of Update objects is returned.
func (bot *Bot) getUpdates(offset uint64) ([]Update, error) {
	params := url.Values{
		"offset":  {fmt.Sprintf("%d", offset)},
		"timeout": {fmt.Sprintf("%d", bot.Timeout/time.Second)},
	}

	response, err := bot.get("getUpdates", params)
	if err != nil {
		return nil, err
	}

	updates := []Update{}
	if err := json.Unmarshal(response.Result, &updates); err != nil {
		return nil, fmt.Errorf("Decode result error (%s)", err.Error())
	}

	return updates, nil
}

// A simple method for testing your bot's auth token.
// Returns basic information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	response, err := bot.get("getMe", url.Values{})
	if err != nil {
		return nil, err
	}

	me := new(User)
	if err := json.Unmarshal(response.Result, me); err != nil {
		return nil, fmt.Errorf("Decode result error (%s)", err.Error())
	} else {
		return me, nil
	}
}

// Use this method to send text messages. On success, the sent Message is returned.
func (bot *Bot) SendMessage(chatId uint64, text string, options *SendMessageOptions) (*Message, error) {
	params := SendMessageParams{
		ChatId: chatId,
		Text:   text,
	}
	if options != nil {
		params.SendMessageOptions = *options
	}

	response, err := bot.post("sendMessage", params)
	if err != nil {
		return nil, err
	}

	message := new(Message)
	if err := json.Unmarshal(response.Result, message); err != nil {
		return nil, fmt.Errorf("Decode result error (%s)", err.Error())
	} else {
		return message, nil
	}
}

// Start getting updates
func (bot *Bot) Start() {
	offset := uint64(0)

	for {
		updates, err := bot.getUpdates(offset + 1)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		for _, update := range updates {
			bot.Updates <- update

			offset = update.UpdateId
		}
	}
}
