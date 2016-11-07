package micha

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/onrik/micha/http"
)

const (
	API_URL      = "https://api.telegram.org/bot%s/%s"
	FILE_API_URL = "https://api.telegram.org/file/bot%s/%s"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "micha ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Set logger
func SetLogger(l *log.Logger) {
	logger = l
}

type Response struct {
	Ok          bool            `json:"ok"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
	Result      json.RawMessage `json:"result"`
}

type Bot struct {
	token   string
	updates chan Update
	Me      User
	Timeout time.Duration
}

// Create new bot instance
func NewBot(token string) (*Bot, error) {
	bot := Bot{
		token:   token,
		updates: make(chan Update),
		Timeout: 25 * time.Second,
	}

	if me, err := bot.GetMe(); err != nil {
		return nil, err
	} else {
		bot.Me = *me
		return &bot, nil
	}
}

// Build url for API method
func (bot *Bot) buildURL(method string) string {
	return fmt.Sprintf(API_URL, bot.token, method)
}

// Decode response result to target object
func (bot *Bot) decodeResponse(data []byte, target interface{}) error {
	response := new(Response)
	if err := json.Unmarshal(data, response); err != nil {
		return fmt.Errorf("Decode response error (%s)", err.Error())
	}

	if !response.Ok {
		return fmt.Errorf("Response status: %d (%s)", response.ErrorCode, response.Description)
	}

	if target == nil {
		// Don't need to decode result
		return nil
	}

	if err := json.Unmarshal(response.Result, target); err != nil {
		return fmt.Errorf("Decode result error (%s)", err.Error())
	} else {
		return nil
	}
}

// Make GET request to Telegram API
func (bot *Bot) get(method string, params url.Values, target interface{}) error {
	response, err := http.Get(bot.buildURL(method) + "?" + params.Encode())
	if err != nil {
		return err
	} else {
		return bot.decodeResponse(response, target)
	}
}

// Make POST request to Telegram API
func (bot *Bot) post(method string, data, target interface{}) error {
	response, err := http.Post(bot.buildURL(method), data)
	if err != nil {
		return err
	} else {
		return bot.decodeResponse(response, target)
	}
}

// Make POST request to Telegram API
func (bot *Bot) postMultipart(method string, file *http.File, params url.Values, target interface{}) error {
	response, err := http.PostMultipart(bot.buildURL(method), file, params)
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
			logger.Printf("Get updates error (%s)\n", err.Error())
			continue
		}

		for _, update := range updates {
			bot.updates <- update

			offset = update.UpdateID
		}
	}
}

// Updates channel
func (bot *Bot) Updates() <-chan Update {
	return bot.updates
}

// A simple method for testing your bot's auth token.
// Returns basic information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	me := new(User)
	err := bot.get("getMe", url.Values{}, me)

	return me, err
}

// Use this method to send text messages.
func (bot *Bot) SendMessage(chatID ChatID, text string, options *SendMessageOptions) (*Message, error) {
	params := sendMessageParams{
		ChatID: chatID,
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
func (bot *Bot) SendPhoto(chatID ChatID, photoID string, options *SendPhotoOptions) (*Message, error) {
	params := newSendPhotoParams(chatID, photoID, options)

	message := new(Message)
	err := bot.post("sendPhoto", params, message)

	return message, err
}

// Send photo file
func (bot *Bot) SendPhotoFile(chatID ChatID, file io.ReadCloser, fileName string, options *SendPhotoOptions) (*Message, error) {
	params := newSendPhotoParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &http.File{
		Source:    file,
		Fieldname: "photo",
		Filename:  fileName,
	}

	message := new(Message)
	err = bot.postMultipart("sendPhoto", f, values, message)

	return message, err
}

// Send exists audio by file_id
func (bot *Bot) SendAudio(chatID ChatID, audioID string, options *SendAudioOptions) (*Message, error) {
	params := newSendAudioParams(chatID, audioID, options)

	message := new(Message)
	err := bot.post("sendAudio", params, message)

	return message, err
}

// Send audio file
func (bot *Bot) SendAudioFile(chatID ChatID, file io.ReadCloser, fileName string, options *SendAudioOptions) (*Message, error) {
	params := newSendAudioParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &http.File{
		Source:    file,
		Fieldname: "audio",
		Filename:  fileName,
	}

	message := new(Message)
	err = bot.postMultipart("sendAudio", f, values, message)

	return message, err
}

// Send exists document by file_id
func (bot *Bot) SendDocument(chatID ChatID, documentID string, options *SendDocumentOptions) (*Message, error) {
	params := newSendDocumentParams(chatID, documentID, options)

	message := new(Message)
	err := bot.post("sendDocument", params, message)

	return message, err
}

// Send file
func (bot *Bot) SendDocumentFile(chatID ChatID, file io.ReadCloser, fileName string, options *SendDocumentOptions) (*Message, error) {
	params := newSendDocumentParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &http.File{
		Source:    file,
		Fieldname: "document",
		Filename:  fileName,
	}

	message := new(Message)
	err = bot.postMultipart("sendDocument", f, values, message)

	return message, err
}

// Send exists sticker by file_id
func (bot *Bot) SendSticker(chatID ChatID, stickerID string, options *SendStickerOptions) (*Message, error) {
	params := newSendStickerParams(chatID, stickerID, options)

	message := new(Message)
	err := bot.post("sendSticker", params, message)

	return message, err
}

// Send .webp sticker file
func (bot *Bot) SendStickerFile(chatID ChatID, file io.ReadCloser, fileName string, options *SendStickerOptions) (*Message, error) {
	params := newSendStickerParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &http.File{
		Source:    file,
		Fieldname: "sticker",
		Filename:  fileName,
	}

	message := new(Message)
	err = bot.postMultipart("sendSticker", f, values, message)

	return message, err
}

// Send exists video by file_id
func (bot *Bot) SendVideo(chatID ChatID, videoID string, options *SendVideoOptions) (*Message, error) {
	params := newSendVideoParams(chatID, videoID, options)

	message := new(Message)
	err := bot.post("sendVideo", params, message)

	return message, err
}

// Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
func (bot *Bot) SendVideoFile(chatID ChatID, file io.ReadCloser, fileName string, options *SendVideoOptions) (*Message, error) {
	params := newSendVideoParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &http.File{
		Source:    file,
		Fieldname: "video",
		Filename:  fileName,
	}

	message := new(Message)
	err = bot.postMultipart("sendVideo", f, values, message)

	return message, err
}

// Send exists voice by file_id
func (bot *Bot) SendVoice(chatID ChatID, voiceID string, options *SendVoiceOptions) (*Message, error) {
	params := newSendVoiceParams(chatID, voiceID, options)

	message := new(Message)
	err := bot.post("sendVoice", params, message)

	return message, err
}

// Use this method to send audio files,
// if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
func (bot *Bot) SendVoiceFile(chatID ChatID, file io.ReadCloser, fileName string, options *SendVoiceOptions) (*Message, error) {
	params := newSendVoiceParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &http.File{
		Source:    file,
		Fieldname: "voice",
		Filename:  fileName,
	}

	message := new(Message)
	err = bot.postMultipart("sendVoice", f, values, message)

	return message, err
}

// Use this method to send point on the map
func (bot *Bot) SendLocation(chatID ChatID, latitude, longitude float64, options *SendLocationOptions) (*Message, error) {
	params := newSendLocationParams(chatID, latitude, longitude, options)

	message := new(Message)
	err := bot.post("sendLocation", params, message)

	return message, err
}

// Use this method to send information about a venue
func (bot *Bot) SendVenue(chatID ChatID, latitude, longitude float64, title, address string, options *SendVenueOptions) (*Message, error) {
	params := newSendVenueParams(chatID, latitude, longitude, title, address, options)

	message := new(Message)
	err := bot.post("sendVenue", params, message)

	return message, err
}

// Use this method to send phone contacts
func (bot *Bot) SendContact(chatID ChatID, phoneNumber, firstName, lastName string, options *SendContactOptions) (*Message, error) {
	params := newSendContactParams(chatID, phoneNumber, firstName, lastName, options)

	message := new(Message)
	err := bot.post("sendContact", params, message)

	return message, err
}

// Use this method to forward messages of any kind.
func (bot *Bot) ForwardMessage(chatID, fromChatID ChatID, messageID int64, disableNotification bool) (*Message, error) {
	params := map[string]interface{}{
		"chat_id":              chatID,
		"from_chat_id":         fromChatID,
		"message_id":           messageID,
		"disable_notification": disableNotification,
	}

	message := new(Message)
	err := bot.post("forwardMessage", params, message)

	return message, err
}

// Use this method when you need to tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
func (bot *Bot) SendChatAction(chatID ChatID, action ChatAction) error {
	params := map[string]interface{}{
		"chat_id": chatID,
		"action":  action,
	}

	return bot.post("sendChatAction", params, nil)
}

// Use this method to get a list of profile pictures for a user.
func (bot *Bot) GetUserProfilePhotos(userID int64, offset, limit *int) (*UserProfilePhotos, error) {
	params := url.Values{
		"user_id": {fmt.Sprintf("%d", userID)},
	}

	if offset != nil {
		params["offset"] = []string{fmt.Sprintf("%d", *offset)}
	}
	if limit != nil {
		params["limit"] = []string{fmt.Sprintf("%d", *limit)}
	}

	profilePhotos := new(UserProfilePhotos)
	err := bot.get("getUserProfilePhotos", params, profilePhotos)

	return profilePhotos, err
}

// Use this method to get basic info about a file and prepare it for downloading.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile again.
func (bot *Bot) GetFile(fileID string) (*File, error) {
	params := url.Values{
		"file_id": {fileID},
	}

	file := new(File)
	err := bot.get("getFile", params, file)

	return file, err
}

// Return absolute url for file downloading by file path
func (bot *Bot) DownloadFileURL(filePath string) string {
	return fmt.Sprintf(FILE_API_URL, bot.token, filePath)
}

// Use this method to edit text messages sent by the bot or via the bot (for inline bots).
func (bot *Bot) EditMessageText(chatID ChatID, messageID int64, inlineMessageID, text string, options *EditMessageTextOptions) (*Message, error) {
	params := editMessageTextParams{
		ChatID:          chatID,
		MessageID:       messageID,
		InlineMessageID: inlineMessageID,
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
func (bot *Bot) EditMessageCaption(chatID ChatID, messageID int64, inlineMessageID string, options *EditMessageCationOptions) (*Message, error) {
	params := editMessageCationParams{
		ChatID:          chatID,
		MessageID:       messageID,
		InlineMessageID: inlineMessageID,
	}
	if options != nil {
		params.EditMessageCationOptions = *options
	}

	message := new(Message)
	err := bot.post("editMessageCaption", params, message)

	return message, err
}

// Use this method to edit only the reply markup of messages sent by the bot or via the bot (for inline bots).
func (bot *Bot) EditMessageReplyMarkup(chatID ChatID, messageID int64, inlineMessageID string, replyMarkup ReplyMarkup) (*Message, error) {
	params := editMessageReplyMarkupParams{
		ChatID:          chatID,
		MessageID:       messageID,
		InlineMessageID: inlineMessageID,
		ReplyMarkup:     replyMarkup,
	}

	message := new(Message)
	err := bot.post("editMessageReplyMarkup", params, message)

	return message, err
}

// Use this method to send answers to an inline query.
// No more than 50 results per query are allowed.
func (bot *Bot) AnswerInlineQuery(inlineQueryID string, results InlineQueryResults, options *AnswerInlineQueryOptions) error {
	params := answerInlineQueryParams{
		InlineQueryID: inlineQueryID,
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
func (bot *Bot) KickChatMember(chatID ChatID, userID int64) error {
	params := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	return bot.post("kickChatMember", params, nil)
}

// Use this method for your bot to leave a group, supergroup or channel
func (bot *Bot) LeaveChat(chatID ChatID) error {
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return bot.post("leaveChat", params, nil)
}

// Use this method to unban a previously kicked user in a supergroup.
// The user will not return to the group automatically, but will be able to join via link, etc.
// The bot must be an administrator in the group for this to work.
func (bot *Bot) UnbanChatMember(chatID ChatID, userID int64) error {
	params := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	return bot.post("unbanChatMember", params, nil)
}

// Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.).
func (bot *Bot) GetChat(chatID ChatID) (*Chat, error) {
	params := url.Values{
		"chat_id": {string(chatID)},
	}

	chat := new(Chat)
	err := bot.get("getChat", params, chat)

	return chat, err
}

// Use this method to get a list of administrators in a chat.
// If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
func (bot *Bot) GetChatAdministrators(chatID ChatID) ([]ChatMember, error) {
	params := url.Values{
		"chat_id": {string(chatID)},
	}

	administrators := []ChatMember{}
	err := bot.get("getChatAdministrators", params, &administrators)

	return administrators, err
}

// Use this method to get the number of members in a chat.
func (bot *Bot) GetChatMembersCount(chatID ChatID) (int, error) {
	params := url.Values{
		"chat_id": {string(chatID)},
	}

	count := 0
	err := bot.get("getChatMembersCount", params, &count)

	return count, err
}

// Use this method to get information about a member of a chat.
func (bot *Bot) GetChatMember(chatID ChatID, userID int64) (*ChatMember, error) {
	params := url.Values{
		"chat_id": {string(chatID)},
		"user_id": {fmt.Sprintf("%d", userID)},
	}

	chatMember := new(ChatMember)
	err := bot.get("getChatMember", params, chatMember)

	return chatMember, err
}

// Use this method to send answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
func (bot *Bot) AnswerCallbackQuery(callbackQueryID string, options *AnswerCallbackQueryOptions) error {
	params := newAnswerCallbackQueryParams(callbackQueryID, options)

	return bot.post("answerCallbackQuery", params, nil)
}

// Use this method to send a game.
func (bot *Bot) SendGame(chatID ChatID, gameShortName string, options *SendGameOptions) (*Message, error) {
	params := sendGameParams{
		ChatID:        chatID,
		GameShortName: gameShortName,
	}

	if options != nil {
		params.SendGameOptions = *options
	}

	message := new(Message)
	err := bot.post("sendGame", params, message)

	return message, err
}

// Use this method to set the score of the specified user in a game.
func (bot *Bot) SetGameScore(userID int64, score int, options *SetGameScoreOptions) (*Message, error) {
	params := setGameScoreParams{
		UserID: userID,
		Score:  score,
	}

	if options != nil {
		params.SetGameScoreOptions = *options
	}

	message := new(Message)
	err := bot.post("setGameScore", params, message)

	return message, err
}

// Use this method to get data for high score tables.
// Will return the score of the specified user and several of his neighbors in a game.
func (bot *Bot) GetGameHighScores(userID int64, options *GetGameHighScoresOptions) ([]GameHighScore, error) {
	params, err := structToValues(options)
	if err != nil {
		return nil, err
	}

	params.Set("user_id", fmt.Sprintf("%d", userID))

	scores := []GameHighScore{}
	err = bot.get("getGameHighScores", params, &scores)

	return scores, err
}
