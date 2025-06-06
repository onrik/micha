package micha

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

const (
	defaultAPIServer = "https://api.telegram.org"
)

type Response struct {
	Ok          bool            `json:"ok"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
	Result      json.RawMessage `json:"result"`
}

// Bot telegram bot
type Bot struct {
	Options
	Me User

	token      string
	updates    chan Update
	offset     uint64
	cancelFunc context.CancelFunc
}

// NewBot - create new bot instance
func NewBot(token string, opts ...Option) (*Bot, error) {
	options := Options{
		limit:      100,
		timeout:    25,
		logger:     slog.Default(),
		apiServer:  defaultAPIServer,
		httpClient: http.DefaultClient,
		ctx:        context.Background(),
	}

	for _, opt := range opts {
		opt(&options)
	}

	bot := Bot{
		Options: options,
		token:   token,
		updates: make(chan Update),
	}
	bot.ctx, bot.cancelFunc = context.WithCancel(options.ctx)

	me, err := bot.GetMe()
	if err != nil {
		return nil, err
	}

	bot.Me = *me

	return &bot, nil
}

// Build url for API method
func (bot *Bot) buildURL(method string) string {
	return bot.Options.apiServer + fmt.Sprintf("/bot%s/%s", bot.token, method)
}

// Decode response result to target object
func (bot *Bot) decodeResponse(data []byte, target interface{}) error {
	response := new(Response)
	if err := json.Unmarshal(data, response); err != nil {
		return fmt.Errorf("decode response error: %w", err)
	}

	if !response.Ok {
		return fmt.Errorf("Error %d (%s)", response.ErrorCode, response.Description)
	}

	if target == nil {
		// Don't need to decode result
		return nil
	}

	if err := json.Unmarshal(response.Result, target); err != nil {
		return fmt.Errorf("decode result error: %w", err)
	}

	return nil
}

// Send GET request to Telegram API
func (bot *Bot) get(method string, params url.Values, target interface{}) error {
	request, err := newGetRequest(bot.ctx, bot.buildURL(method), params)
	if err != nil {
		return err
	}

	response, err := bot.httpClient.Do(request)
	if err != nil {
		return err
	}

	body, err := handleResponse(response)
	if err != nil {
		return err
	}

	return bot.decodeResponse(body, target)
}

// Send POST request to Telegram API
func (bot *Bot) post(method string, data, target interface{}) error {
	request, err := newPostRequest(bot.ctx, bot.buildURL(method), data)
	if err != nil {
		return err
	}
	response, err := bot.httpClient.Do(request)
	if err != nil {
		return err
	}

	body, err := handleResponse(response)
	if err != nil {
		return err
	}

	return bot.decodeResponse(body, target)
}

// Send POST multipart request to Telegram API
func (bot *Bot) postMultipart(method string, file *fileField, params url.Values, target interface{}) error {
	request, err := newMultipartRequest(bot.ctx, bot.buildURL(method), file, params)
	if err != nil {
		return err
	}
	response, err := bot.httpClient.Do(request)
	if err != nil {
		return err
	}

	body, err := handleResponse(response)
	if err != nil {
		return err
	}

	return bot.decodeResponse(body, target)
}

// Use this method to receive incoming updates using long polling.
// An Array of Update objects is returned.
func (bot *Bot) getUpdates(offset uint64, allowedUpdates ...string) ([]Update, error) {
	params := url.Values{
		"limit":   {fmt.Sprintf("%d", bot.limit)},
		"offset":  {fmt.Sprintf("%d", offset)},
		"timeout": {fmt.Sprintf("%d", bot.timeout)},
	}

	if len(allowedUpdates) > 0 {
		params["allowed_updates"] = allowedUpdates
	}

	updates := []Update{}
	err := bot.get("getUpdates", params, &updates)

	return updates, err
}

// Start getting updates
func (bot *Bot) Start(allowedUpdates ...string) {
	for {
		updates, err := bot.getUpdates(bot.offset+1, allowedUpdates...)
		if err != nil {
			bot.logger.ErrorContext(bot.ctx, "Get updates error", "error", err)
			httpErr := HTTPError{}
			if errors.As(err, &httpErr) && httpErr.StatusCode == http.StatusConflict {
				bot.cancelFunc()
			}
		}

		for _, update := range updates {
			bot.updates <- update
			bot.offset = update.UpdateID
		}

		select {
		case <-bot.ctx.Done():
			close(bot.updates)
			return
		default:
		}
	}
}

// Stop getting updates
func (bot *Bot) Stop() {
	bot.cancelFunc()
}

// Updates channel
func (bot *Bot) Updates() <-chan Update {
	return bot.updates
}

func (bot *Bot) GetWebhookInfo() (*WebhookInfo, error) {
	webhookInfo := new(WebhookInfo)
	err := bot.get("getWebhookInfo", url.Values{}, webhookInfo)

	return webhookInfo, err
}

func (bot *Bot) SetWebhook(webhookURL string, options *SetWebhookOptions) error {
	var file *fileField
	params := url.Values{
		"url": {webhookURL},
	}
	if options != nil {
		if options.MaxConnections > 0 {
			params.Set("max_connections", fmt.Sprintf("%d", options.MaxConnections))
		}
		if len(options.AllowedUpdates) > 0 {
			params["allowed_updates"] = options.AllowedUpdates
		}
		if len(options.Certificate) > 0 {
			file = &fileField{
				Source:    bytes.NewBuffer(options.Certificate),
				Fieldname: "certificate",
				Filename:  "certificate",
			}
		}
	}

	return bot.postMultipart("setWebhook", file, params, nil)
}

func (bot *Bot) DeleteWebhook() error {
	return bot.post("deleteWebhook", nil, nil)
}

// Logout
// Use this method to log out from the cloud Bot API server before launching the bot locally.
// You must log out the bot before running it locally,
// otherwise there is no guarantee that the bot will receive updates.
// After a successful call, you can immediately log in on a local server,
// but will not be able to log in back to the cloud Bot API server for 10 minutes.
func (bot *Bot) Logout() error {
	url := defaultAPIServer + fmt.Sprintf("/bot%s/logOut", bot.token)
	request, err := newGetRequest(bot.ctx, url, nil)
	if err != nil {
		return err
	}

	response, err := bot.httpClient.Do(request)
	if err != nil {
		return err
	}

	_, err = handleResponse(response)
	return err
}

// A simple method for testing your bot's auth token.
// Returns basic information about the bot in form of a User object.
func (bot *Bot) GetMe() (*User, error) {
	me := new(User)
	err := bot.get("getMe", nil, me)

	return me, err
}

// Raw - send any method and return raw response
func (bot *Bot) Raw(method string, data any) ([]byte, error) {
	request, err := newPostRequest(bot.ctx, bot.buildURL(method), data)
	if err != nil {
		return nil, err
	}

	response, err := bot.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return handleResponse(response)
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
func (bot *Bot) SendPhotoFile(chatID ChatID, file io.Reader, fileName string, options *SendPhotoOptions) (*Message, error) {
	params := newSendPhotoParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &fileField{
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
func (bot *Bot) SendAudioFile(chatID ChatID, file io.Reader, fileName string, options *SendAudioOptions) (*Message, error) {
	params := newSendAudioParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &fileField{
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
func (bot *Bot) SendDocumentFile(chatID ChatID, file io.Reader, fileName string, options *SendDocumentOptions) (*Message, error) {
	params := newSendDocumentParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &fileField{
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
func (bot *Bot) SendStickerFile(chatID ChatID, file io.Reader, fileName string, options *SendStickerOptions) (*Message, error) {
	params := newSendStickerParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &fileField{
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
func (bot *Bot) SendVideoFile(chatID ChatID, file io.Reader, fileName string, options *SendVideoOptions) (*Message, error) {
	params := newSendVideoParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &fileField{
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
func (bot *Bot) SendVoiceFile(chatID ChatID, file io.Reader, fileName string, options *SendVoiceOptions) (*Message, error) {
	params := newSendVoiceParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &fileField{
		Source:    file,
		Fieldname: "voice",
		Filename:  fileName,
	}

	message := new(Message)
	err = bot.postMultipart("sendVoice", f, values, message)

	return message, err
}

// Send exists video note by file_id
func (bot *Bot) SendVideoNote(chatID ChatID, videoNoteID string, options *SendVideoNoteOptions) (*Message, error) {
	params := newSendVideoNoteParams(chatID, videoNoteID, options)

	message := new(Message)
	err := bot.post("sendVideoNote", params, message)

	return message, err
}

// Use this method to send video messages
func (bot *Bot) SendVideoNoteFile(chatID ChatID, file io.Reader, fileName string, options *SendVideoNoteOptions) (*Message, error) {
	params := newSendVideoNoteParams(chatID, "", options)
	values, err := structToValues(params)
	if err != nil {
		return nil, err
	}

	f := &fileField{
		Source:    file,
		Fieldname: "video_note",
		Filename:  fileName,
	}

	message := new(Message)
	err = bot.postMultipart("sendVideoNote", f, values, message)

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
	return bot.Options.apiServer + fmt.Sprintf("/file/bot%s/%s", bot.token, filePath)
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

// Use this method to delete a message.
// A message can only be deleted if it was sent less than 48 hours ago.
// Any such recently sent outgoing message may be deleted.
// Additionally, if the bot is an administrator in a group chat, it can delete any message.
// If the bot is an administrator in a supergroup, it can delete messages from any other user and service messages about people joining or leaving the group (other types of service messages may only be removed by the group creator). In channels, bots can only remove their own messages.
func (bot *Bot) DeleteMessage(chatID ChatID, messageID int64) (bool, error) {
	params := map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
	}

	var success bool
	err := bot.post("deleteMessage", params, &success)

	return success, err
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
