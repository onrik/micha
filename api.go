package micha

type CommonSendParams struct {
	DisableNotification bool  `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64 `json:"reply_to_message_id,omitempty"`
}

type SendMessageOptions struct {
	CommonSendParams
	ParseMode             ParseMode    `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool         `json:"disable_web_page_preview",omitempty`
	ReplyMarkup           *ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendMessageParams struct {
	SendMessageOptions
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type ReplyMarkup struct {
	ForceReply         bool                     `json:"force_reply,omitempty"`
	CustomKeyboard     [][]string               `json:"keyboard,omitempty"`
	ResizeKeyboard     bool                     `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard    bool                     `json:"one_time_keyboard,omitempty"`
	HideCustomKeyboard bool                     `json:"hide_keyboard,omitempty"`
	Selective          bool                     `json:"selective,omitempty"`
	InlineKeyboard     [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

type EditMessageTextOptions struct {
	ParseMode             ParseMode    `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool         `json:"disable_web_page_preview",omitempty`
	ReplyMarkup           *ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageTextParams struct {
	ChatId          int64  `json:"chat_id,omitempty"`
	MessageId       int64  `json:"message_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	Text            string `json:"text"`
	EditMessageTextOptions
}

type EditMessageCationOptions struct {
	Caption     string       `json:"captino,omitempty"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageCationParams struct {
	ChatId          int64  `json:"chat_id,omitempty"`
	MessageId       int64  `json:"message_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	EditMessageCationOptions
}

type EditMessageReplyMarkupParams struct {
	ChatId          int64        `json:"chat_id,omitempty"`
	MessageId       int64        `json:"message_id,omitempty"`
	InlineMessageId string       `json:"inline_message_id,omitempty"`
	ReplyMarkup     *ReplyMarkup `json:"reply_markup,omitempty"`
}
