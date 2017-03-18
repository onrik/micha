package micha

type Options struct {
	Limit   int
	Timeout int
	Logger  Logger
}

type Option func(*Options)

// WithLimit - set getUpdates limit
// Values between 1—100 are accepted. Defaults to 100.
func WithLimit(limit int) Option {
	return func(o *Options) {
		o.Limit = limit
	}
}

// WithTimeout - set timeout in seconds for getUpdates long polling
// Defaults to 25
func WithTimeout(timeout int) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

// WithLogger - set logger
func WithLogger(logger Logger) Option {
	return func(o *Options) {
		o.Logger = logger
	}
}

// Send message optional params
type SendMessageOptions struct {
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int64       `json:"reply_to_message_id,omitempty"`
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send photo optional params
type SendPhotoOptions struct {
	Caption             string      `json:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send audio optional params
type SendAudioOptions struct {
	Duration            int         `json:"duration,omitempty"`
	Performer           string      `json:"performer,omitempty"`
	Title               string      `json:"title,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send document optional params
type SendDocumentOptions struct {
	Caption             string      `json:"caption,omitempty"`
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

// Send video optional params
type SendVideoOptions struct {
	Duration            int         `json:"duration,omitempty"`
	Width               int         `json:"width,omitempty"`
	Height              int         `json:"height,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send voice optional params
type SendVoiceOptions struct {
	Duration            int         `json:"duration,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send location optional params
type SendLocationOptions struct {
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send venue optional params
type SendVenueOptions struct {
	FoursquareID        string      `json:"foursquare_id,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send contact optional params
type SendContactOptions struct {
	DisableNotification bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID    int64       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
}

// Send game optional params
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
