package micha

import (
	"context"
	"strings"
)

type Options struct {
	limit      int
	timeout    int
	logger     Logger
	apiServer  string
	httpClient HttpClient
	ctx        context.Context
}

type Option func(*Options)

// WithLimit - set getUpdates limit
// Values between 1—100 are accepted. Defaults to 100.
func WithLimit(limit int) Option {
	return func(o *Options) {
		o.limit = limit
	}
}

// WithTimeout - set timeout in seconds for getUpdates long polling
// Defaults to 25
func WithTimeout(timeout int) Option {
	return func(o *Options) {
		o.timeout = timeout
	}
}

// WithLogger - set logger
func WithLogger(logger Logger) Option {
	return func(o *Options) {
		o.logger = logger
	}
}

// WithHttpClient - set custom http client
func WithHttpClient(httpClient HttpClient) Option {
	return func(o *Options) {
		o.httpClient = httpClient
	}
}

// WithAPIServer - set custom api server (https://github.com/tdlib/telegram-bot-api)
func WithAPIServer(url string) Option {
	return func(o *Options) {
		o.apiServer = strings.TrimSuffix(url, "/")
	}
}

// WithAPIServer - set custom context
func WithCtx(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

// SendMessageOptions optional params SendMessage method
type SendMessageOptions struct {
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendPhotoOptions optional params SendPhoto method
type SendPhotoOptions struct {
	Caption             string      `json:"caption,omitempty"`
	ParseMode           ParseMode   `json:"parse_mode,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendAudioOptions optional params SendAudio method
type SendAudioOptions struct {
	Caption             string      `json:"caption,omitempty"`
	ParseMode           ParseMode   `json:"parse_mode,omitempty"`
	Duration            int         `json:"duration,omitempty"`
	Performer           string      `json:"performer,omitempty"`
	Title               string      `json:"title,omitempty"`
	Thumb               string      `json:"thumb,omitempty"` // TODO add thumb as file
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ProtectContent      bool        `json:"protect_content,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendDocumentOptions optional params SendDocument method
type SendDocumentOptions struct {
	Thumb               string      `json:"thumb,omitempty"` // TODO add thumb as file
	Caption             string      `json:"caption,omitempty"`
	ParseMode           ParseMode   `json:"parse_mode,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send sticker optional params
type SendStickerOptions struct {
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVideoOptions video optional params SendVideo method
type SendVideoOptions struct {
	Duration            int         `json:"duration,omitempty"`
	Width               int         `json:"width,omitempty"`
	Height              int         `json:"height,omitempty"`
	Thumb               string      `json:"thumb,omitempty"` // TODO add thumb as file
	Caption             string      `json:"caption,omitempty"`
	ParseMode           ParseMode   `json:"parse_mode,omitempty"`
	SupportsStreaming   bool        `json:"supports_streaming,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVoiceOptions optional params for SendVoice method
type SendVoiceOptions struct {
	Caption             string      `json:"caption,omitempty"`
	ParseMode           ParseMode   `json:"parse_mode,omitempty"`
	Duration            int         `json:"duration,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVideoNoteOptions optional params for SendVideoNote method
type SendVideoNoteOptions struct {
	Duration            int         `json:"duration,omitempty"`
	Length              int         `json:"length,omitempty"`
	Thumb               string      `json:"thumb,omitempty"` // TODO add thumb as file
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendLocationOptions optional params for SendLocation method
type SendLocationOptions struct {
	LivePeriod          int         `json:"live_period,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVenueOptions optional params for SendVenue method
type SendVenueOptions struct {
	FoursquareID        string      `json:"foursquare_id,omitempty"`
	FoursquareType      string      `json:"foursquare_type,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendContactOptions optional params for SendContact method
type SendContactOptions struct {
	VCard               string      `json:"vcard,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendGameOptions optional params for SendGame method
type SendGameOptions struct {
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Set game score optional params
type SetGameScoreOptions struct {
	ChatID             ChatID `json:"chat_id,omitempty"`
	MessageID          int64  `json:"message_id,omitempty"`
	InlineMessageID    string `json:"inline_message_id,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	Force              bool   `json:"force,omitempty"`
}

// Get game high scopres optional params
type GetGameHighScoresOptions struct {
	ChatID          ChatID `json:"chat_id,omitempty"`
	MessageID       int64  `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// Edit message text optional params
type EditMessageTextOptions struct {
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

// Edit message caption optional params
type EditMessageCationOptions struct {
	Caption     string      `json:"caption,omitempty"`
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Answer callback query optional params
type AnswerCallbackQueryOptions struct {
	Text      string `json:"text,omitempty"`
	ShowAlert bool   `json:"show_alert,omitempty"`
	URL       string `json:"url,omitempty"`
	CacheTime int    `json:"cache_time,omitempty"`
}

// Answer inline query optional params
type AnswerInlineQueryOptions struct {
	CacheTime         int    `json:"cache_time,omitempty"`
	IsPersonal        bool   `json:"is_personal,omitempty"`
	NextOffset        string `json:"next_offset,omitempty"`
	SwitchPmText      string `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string `json:"switch_pm_parameter,omitempty"`
}

// Set webhook query optional params
type SetWebhookOptions struct {
	Certificate    []byte   `json:"certificate,omitempty"`
	MaxConnections int      `json:"max_connections,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}
