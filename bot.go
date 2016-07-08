package micha

import (
	"encoding/json"
	"fmt"
	"io"
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

// Decode response result to target object
func (bot *Bot) decodeResponse(data []byte, target interface{}) error {
	apiResponse := new(ApiResponse)
	if err := json.Unmarshal(data, apiResponse); err != nil {
		return fmt.Errorf("Decode response error (%s)", err.Error())
	}

	if !apiResponse.Ok {
		return fmt.Errorf("Response status: %d (%s)", apiResponse.ErrorCode, apiResponse.Description)
	}

	if target == nil {
		// Don't need to decode result
		return nil
	}

	if err := json.Unmarshal(apiResponse.Result, target); err != nil {
		return fmt.Errorf("Decode result error (%s)", err.Error())
	} else {
		return nil
	}
}

// Make GET request to Telegram API
func (bot *Bot) get(method string, params url.Values, target interface{}) error {
	response, err := get(bot.buildUrl(method) + "?" + params.Encode())
	if err != nil {
		return err
	} else {
		return bot.decodeResponse(response, target)
	}
}

// Make POST request to Telegram API
func (bot *Bot) post(method string, data, target interface{}) error {
	response, err := post(bot.buildUrl(method), data)
	if err != nil {
		return err
	} else {
		return bot.decodeResponse(response, target)
	}
}

// Make POST request to Telegram API
func (bot *Bot) postMultipart(method string, file *FileToSend, params url.Values, target interface{}) error {
	response, err := postMultipart(bot.buildUrl(method), file, params)
	if err != nil {
		return err
	} else {
		return bot.decodeResponse(response, target)
	}
}

// Use this method to receive incoming updates using long polling.
// An Array of Update objects is returned.
func (bot *Bot) getUpdates(offset uint64) ([]Update, error) {
	params := url.Values{
		"offset":  {fmt.Sprintf("%d", offset)},
		"timeout": {fmt.Sprintf("%d", bot.Timeout/time.Second)},
	}

	updates := []Update{}
	err := bot.get("getUpdates", params, &updates)

	return updates, err
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

// A simple method for testing your bot's auth token.
// Returns basic information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	me := new(User)
	err := bot.get("getMe", url.Values{}, me)

	return me, err
}

// Use this method to send text messages.
func (bot *Bot) SendMessage(chatId int64, text string, options *SendMessageOptions) (*Message, error) {
	params := SendMessageParams{
		ChatId: chatId,
		Text:   text,
	}
	if options != nil {
		params.SendMessageOptions = *options
	}

	message := new(Message)
	err := bot.post("sendMessage", params, message)

	return message, err
}

// Send exists photo by file_id
func (bot *Bot) SendPhoto(chatId int64, photoId string, options *SendPhotoOptions) (*Message, error) {
	params := NewSendPhotoParams(chatId, photoId, options)

	message := new(Message)
	err := bot.post("sendPhoto", params, message)

	return message, err
}

// Send photo file
func (bot *Bot) SendPhotoFile(chatId int64, file io.ReadCloser, options *SendPhotoOptions) (*Message, error) {
	params := NewSendPhotoParams(chatId, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	fileToSend := &FileToSend{
		File:      file,
		Fieldname: "photo",
		Filename:  "photo.png",
	}

	message := new(Message)
	err = bot.postMultipart("sendPhoto", fileToSend, values, message)

	return message, err
}

// Send exists audio by file_id
func (bot *Bot) SendAudio(chatId int64, audioId string, options *SendAudioOptions) (*Message, error) {
	params := NewSendAudioParams(chatId, audioId, options)

	message := new(Message)
	err := bot.post("sendAudio", params, message)

	return message, err
}

// Send audio file
func (bot *Bot) SendAudioFile(chatId int64, file io.ReadCloser, options *SendAudioOptions) (*Message, error) {
	params := NewSendAudioParams(chatId, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	fileToSend := &FileToSend{
		File:      file,
		Fieldname: "audio",
		Filename:  "audio.mp3",
	}

	message := new(Message)
	err = bot.postMultipart("sendAudio", fileToSend, values, message)

	return message, err
}

// Send exists document by file_id
func (bot *Bot) SendDocument(chatId int64, documentId string, options *SendDocumentOptions) (*Message, error) {
	params := NewSendDocumentParams(chatId, documentId, options)

	message := new(Message)
	err := bot.post("sendDocument", params, message)

	return message, err
}

// Send file
func (bot *Bot) SendDocumentFile(chatId int64, documentName string, file io.ReadCloser, options *SendDocumentOptions) (*Message, error) {
	params := NewSendDocumentParams(chatId, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	fileToSend := &FileToSend{
		File:      file,
		Fieldname: "document",
		Filename:  documentName,
	}

	message := new(Message)
	err = bot.postMultipart("sendDocument", fileToSend, values, message)

	return message, err
}

// Send exists sticker by file_id
func (bot *Bot) SendSticker(chatId int64, stickerId string, options *SendStickerOptions) (*Message, error) {
	params := NewSendStickerParams(chatId, stickerId, options)

	message := new(Message)
	err := bot.post("sendSticker", params, message)

	return message, err
}

// Send .webp sticker file
func (bot *Bot) SendStickerFile(chatId int64, file io.ReadCloser, options *SendStickerOptions) (*Message, error) {
	params := NewSendStickerParams(chatId, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	fileToSend := &FileToSend{
		File:      file,
		Fieldname: "sticker",
		Filename:  "sticker.webp",
	}

	message := new(Message)
	err = bot.postMultipart("sendSticker", fileToSend, values, message)

	return message, err
}

// Send exists video by file_id
func (bot *Bot) SendVideo(chatId int64, videoId string, options *SendVideoOptions) (*Message, error) {
	params := NewSendVideoParams(chatId, videoId, options)

	message := new(Message)
	err := bot.post("sendVideo", params, message)

	return message, err
}

// Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
func (bot *Bot) SendVideoFile(chatId int64, file io.ReadCloser, options *SendVideoOptions) (*Message, error) {
	params := NewSendVideoParams(chatId, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	fileToSend := &FileToSend{
		File:      file,
		Fieldname: "video",
		Filename:  "video.mp4",
	}

	message := new(Message)
	err = bot.postMultipart("sendVideo", fileToSend, values, message)

	return message, err
}

// Send exists voice by file_id
func (bot *Bot) SendVoice(chatId int64, voiceId string, options *SendVoiceOptions) (*Message, error) {
	params := NewSendVoiceParams(chatId, voiceId, options)

	message := new(Message)
	err := bot.post("sendVoice", params, message)

	return message, err
}

// Use this method to send audio files,
// if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
func (bot *Bot) SendVoiceFile(chatId int64, file io.ReadCloser, options *SendVoiceOptions) (*Message, error) {
	params := NewSendVoiceParams(chatId, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	fileToSend := &FileToSend{
		File:      file,
		Fieldname: "voice",
		Filename:  "voice.ogg",
	}

	message := new(Message)
	err = bot.postMultipart("sendVoice", fileToSend, values, message)

	return message, err
}

// Use this method to send point on the map
func (bot *Bot) SendLocation(chatId int64, latitude, longitude float64, options *SendLocationOptions) (*Message, error) {
	params := NewSendLocationParams(chatId, latitude, longitude, options)

	message := new(Message)
	err := bot.post("sendLocation", params, message)

	return message, err
}

// Use this method to send information about a venue
func (bot *Bot) SendVenue(chatId int64, latitude, longitude float64, title, address string, options *SendVenueOptions) (*Message, error) {
	params := NewSendVenueParams(chatId, latitude, longitude, title, address, options)

	message := new(Message)
	err := bot.post("sendVenue", params, message)

	return message, err
}

// Use this method to send phone contacts
func (bot *Bot) SendContact(chatId int64, phoneNumber, firstName, lastName string, options *SendContactOptions) (*Message, error) {
	params := NewSendContactParams(chatId, phoneNumber, firstName, lastName, options)

	message := new(Message)
	err := bot.post("sendContact", params, message)

	return message, err
}

// Use this method to forward messages of any kind.
func (bot *Bot) ForwardMessage(chatId, fromChatId, messageId int64, disableNotification bool) (*Message, error) {
	params := map[string]interface{}{
		"chat_id":              chatId,
		"from_chat_id":         fromChatId,
		"message_id":           messageId,
		"disable_notification": disableNotification,
	}

	message := new(Message)
	err := bot.post("forwardMessage", params, message)

	return message, err
}

// Use this method when you need to tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
func (bot *Bot) SendChatAction(chatId int64, action ChatAction) error {
	params := map[string]interface{}{
		"chat_id": chatId,
		"action":  action,
	}

	return bot.post("sendChatAction", params, nil)
}

// Use this method to edit text messages sent by the bot or via the bot (for inline bots).
func (bot *Bot) EditMessageText(chatId, messageId int64, inlineMessageId, text string, options *EditMessageTextOptions) (*Message, error) {
	params := EditMessageTextParams{
		ChatId:          chatId,
		MessageId:       messageId,
		InlineMessageId: inlineMessageId,
		Text:            text,
	}
	if options != nil {
		params.EditMessageTextOptions = *options
	}

	message := new(Message)
	err := bot.post("editMessageText", params, message)

	return message, err
}

// Use this method to edit captions of messages sent by the bot or via the bot (for inline bots).
func (bot *Bot) EditMessageCaption(chatId, messageId int64, inlineMessageId string, options *EditMessageCationOptions) (*Message, error) {
	params := EditMessageCationParams{
		ChatId:          chatId,
		MessageId:       messageId,
		InlineMessageId: inlineMessageId,
	}
	if options != nil {
		params.EditMessageCationOptions = *options
	}

	message := new(Message)
	err := bot.post("editMessageCaption", params, message)

	return message, err
}

// Use this method to edit only the reply markup of messages sent by the bot or via the bot (for inline bots).
func (bot *Bot) EditMessageReplyMarkup(chatId, messageId int64, inlineMessageId string, replyMarkup ReplyMarkup) (*Message, error) {
	params := EditMessageReplyMarkupParams{
		ChatId:          chatId,
		MessageId:       messageId,
		InlineMessageId: inlineMessageId,
		ReplyMarkup:     replyMarkup,
	}

	message := new(Message)
	err := bot.post("editMessageReplyMarkup", params, message)

	return message, err
}

// Use this method to send answers to an inline query.
// No more than 50 results per query are allowed.
func (bot *Bot) AnswerInlineQuery(inlineQueryId string, results InlineQueryResults, options *AnswerInlineQueryOptions) error {
	params := AnswerInlineQueryParams{
		InlineQueryId: inlineQueryId,
		Results:       results,
	}
	if options != nil {
		params.AnswerInlineQueryOptions = *options
	}

	return bot.post("answerInlineQuery", params, nil)
}

// Use this method to kick a user from a group or a supergroup.
// In the case of supergroups, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first.
// The bot must be an administrator in the group for this to work.
func (bot *Bot) KickChatMember(chatId, userId int64) error {
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}

	return bot.post("kickChatMember", params, nil)
}

// Use this method for your bot to leave a group, supergroup or channel
func (bot *Bot) LeaveChat(chatId int64) error {
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return bot.post("leaveChat", params, nil)
}

// Use this method to unban a previously kicked user in a supergroup.
// The user will not return to the group automatically, but will be able to join via link, etc.
// The bot must be an administrator in the group for this to work.
func (bot *Bot) UnbanChatMember(chatId, userId int64) error {
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}

	return bot.post("unbanChatMember", params, nil)
}

// Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.).
func (bot *Bot) GetChat(chatId int64) (*Chat, error) {
	params := url.Values{
		"chat_id": {fmt.Sprintf("%d", chatId)},
	}

	chat := new(Chat)
	err := bot.get("getChat", params, chat)

	return chat, err
}

// Use this method to get a list of administrators in a chat.
// If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(chatId int64) ([]ChatMember, error) {
	params := url.Values{
		"chat_id": {fmt.Sprintf("%d", chatId)},
	}

	administrators := []ChatMember{}
	err := bot.get("getChatAdministrators", params, administrators)

	return administrators, err
}

// Use this method to get the number of members in a chat.
func (bot *Bot) GetChatMembersCount(chatId int64) (int, error) {
	params := url.Values{
		"chat_id": {fmt.Sprintf("%d", chatId)},
	}

	count := 0
	err := bot.get("getChatMembersCount", params, &count)

	return count, err
}

// Use this method to get information about a member of a chat.
func (bot *Bot) GetChatMember(chatId, userId int64) (*ChatMember, error) {
	params := url.Values{
		"chat_id": {fmt.Sprintf("%d", chatId)},
		"user_id": {fmt.Sprintf("%d", userId)},
	}

	chatMember := new(ChatMember)
	err := bot.get("getChatMember", params, chatMember)

	return chatMember, err
}
