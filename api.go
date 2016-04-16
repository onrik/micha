package micha

type CommonSendParams struct {
	DisableNotification bool   `json:"disable_notification"`
	ReplyToMessageId    uint64 `json:"reply_to_message_id,omitempty"`
}

type SendMessageOptions struct {
	CommonSendParams
	ParseMode             ParseMode   `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview"`
	ReplyMarkup           ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendMessageParams struct {
	SendMessageOptions
	ChatId uint64 `json:"chat_id"`
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
