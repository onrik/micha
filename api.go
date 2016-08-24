package micha

type SendMessageOptions struct {
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId      int64       `json:"reply_to_message_id,omitempty"`
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendMessageParams struct {
	SendMessageOptions
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// Send photo request params
type SendPhotoOptions struct {
	Caption             string      `json:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendPhotoParams struct {
	ChatId int64  `json:"chat_id"`
	Photo  string `json:"photo,omitempty"`
	SendPhotoOptions
}

func NewSendPhotoParams(chatId int64, photo string, options *SendPhotoOptions) *SendPhotoParams {
	params := &SendPhotoParams{
		ChatId: chatId,
		Photo:  photo,
	}

	if options != nil {
		params.SendPhotoOptions = *options
	}

	return params
}

// Send audio request params
type SendAudioOptions struct {
	Duration            int         `json:"duration,omitempty"`
	Performer           string      `json:"performer,omitempty"`
	Title               string      `json:"title,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendAudioParams struct {
	ChatId int64  `json:"chat_id"`
	Audio  string `json:"audio,omitempty"`
	SendAudioOptions
}

func NewSendAudioParams(chatId int64, audio string, options *SendAudioOptions) *SendAudioParams {
	params := &SendAudioParams{
		ChatId: chatId,
		Audio:  audio,
	}

	if options != nil {
		params.SendAudioOptions = *options
	}

	return params
}

// Send document request params
type SendDocumentOptions struct {
	Caption             string      `json:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendDocumentParams struct {
	ChatId   int64  `json:"chat_id"`
	Document string `json:"document,omitempty"`
	SendDocumentOptions
}

func NewSendDocumentParams(chatId int64, document string, options *SendDocumentOptions) *SendDocumentParams {
	params := &SendDocumentParams{
		ChatId:   chatId,
		Document: document,
	}

	if options != nil {
		params.SendDocumentOptions = *options
	}

	return params
}

// Send sticker request params
type SendStickerOptions struct {
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendStickerParams struct {
	ChatId  int64  `json:"chat_id"`
	Sticker string `json:"sticker,omitempty"`
	SendStickerOptions
}

func NewSendStickerParams(chatId int64, sticker string, options *SendStickerOptions) *SendStickerParams {
	params := &SendStickerParams{
		ChatId:  chatId,
		Sticker: sticker,
	}

	if options != nil {
		params.SendStickerOptions = *options
	}

	return params
}

// Send video request params
type SendVideoOptions struct {
	Duration            int         `json:"duration,omitempty"`
	Width               int         `json:"width,omitempty"`
	Height              int         `json:"height,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendVideoParams struct {
	ChatId int64  `json:"chat_id"`
	Video  string `json:"video,omitempty"`
	SendVideoOptions
}

func NewSendVideoParams(chatId int64, video string, options *SendVideoOptions) *SendVideoParams {
	params := &SendVideoParams{
		ChatId: chatId,
		Video:  video,
	}

	if options != nil {
		params.SendVideoOptions = *options
	}

	return params
}

// Send voice request params
type SendVoiceOptions struct {
	Duration            int         `json:"duration,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendVoiceParams struct {
	ChatId int64  `json:"chat_id"`
	Voice  string `json:"voice,omitempty"`
	SendVoiceOptions
}

func NewSendVoiceParams(chatId int64, voice string, options *SendVoiceOptions) *SendVoiceParams {
	params := &SendVoiceParams{
		ChatId: chatId,
		Voice:  voice,
	}

	if options != nil {
		params.SendVoiceOptions = *options
	}

	return params
}

// Send location request params
type SendLocationOptions struct {
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendLocationParams struct {
	ChatId    int64   `json:"chat_id"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	SendLocationOptions
}

func NewSendLocationParams(chatId int64, latitude, longitude float64, options *SendLocationOptions) *SendLocationParams {
	params := &SendLocationParams{
		ChatId:    chatId,
		Latitude:  latitude,
		Longitude: longitude,
	}

	if options != nil {
		params.SendLocationOptions = *options
	}

	return params
}

// Send venue request params
type SendVenueOptions struct {
	FoursquareId        string      `json:"foursquare_id,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendVenueParams struct {
	ChatId    int64   `json:"chat_id"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Title     string  `json:"title,omitempty"`
	Address   string  `json:"address,omitempty"`
	SendVenueOptions
}

func NewSendVenueParams(chatId int64, latitude, longitude float64, title, address string, options *SendVenueOptions) *SendVenueParams {
	params := &SendVenueParams{
		ChatId:    chatId,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}

	if options != nil {
		params.SendVenueOptions = *options
	}

	return params
}

// Send contact request params
type SendContactOptions struct {
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendContactParams struct {
	ChatId      int64  `json:"chat_id"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	SendContactOptions
}

func NewSendContactParams(chatId int64, phoneNumber, firstName, lastName string, options *SendContactOptions) *SendContactParams {
	params := &SendContactParams{
		ChatId:      chatId,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
		LastName:    lastName,
	}

	if options != nil {
		params.SendContactOptions = *options
	}

	return params
}

// answerCallbackQuery request params
type AnswerCallbackQueryOptions struct {
	Text      string `json:"text,omitempty"`
	ShowAlert bool   `json:"show_alert,omitempty"`
}

type AnswerCallbackQueryParams struct {
	CallbackQueryId string `json:"callback_query_id"`
	AnswerCallbackQueryOptions
}

func NewAnswerCallbackQueryParams(callbackQueryId string, options *AnswerCallbackQueryOptions) *AnswerCallbackQueryParams {
	params := &AnswerCallbackQueryParams{
		CallbackQueryId: callbackQueryId,
	}

	if options != nil {
		params.AnswerCallbackQueryOptions = *options
	}

	return params
}

type EditMessageTextOptions struct {
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageTextParams struct {
	ChatId          int64  `json:"chat_id,omitempty"`
	MessageId       int64  `json:"message_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	Text            string `json:"text"`
	EditMessageTextOptions
}

type EditMessageCationOptions struct {
	Caption     string      `json:"caption,omitempty"`
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageCationParams struct {
	ChatId          int64  `json:"chat_id,omitempty"`
	MessageId       int64  `json:"message_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	EditMessageCationOptions
}

type EditMessageReplyMarkupParams struct {
	ChatId          int64       `json:"chat_id,omitempty"`
	MessageId       int64       `json:"message_id,omitempty"`
	InlineMessageId string      `json:"inline_message_id,omitempty"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`
}

type AnswerInlineQueryOptions struct {
	CacheTime         int    `json:"cache_time,omitempty"`
	IsPersonal        bool   `json:"is_personal,omitempty"`
	NextOffset        string `json:"next_offset,omitempty"`
	SwitchPmText      string `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string `json:"switch_pm_parameter,omitempty"`
}

type AnswerInlineQueryParams struct {
	InlineQueryId string             `json:"inline_query_id"`
	Results       InlineQueryResults `json:"results"`
	AnswerInlineQueryOptions
}
