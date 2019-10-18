package micha

const (
	PARSE_MODE_DEFAULT  ParseMode = ""
	PARSE_MODE_HTML     ParseMode = "HTML"
	PARSE_MODE_MARKDOWN ParseMode = "Markdown"

	CHAT_TYPE_PRIVATE    ChatType = "private"
	CHAT_TYPE_GROUP      ChatType = "group"
	CHAT_TYPE_SUPERGROUP ChatType = "supergroup"
	CHAT_TYPE_CHANNEL    ChatType = "channel"

	CHAT_ACTION_TYPING            ChatAction = "typing"
	CHAT_ACTION_UPLOAD_PHOTO      ChatAction = "upload_photo"
	CHAT_ACTION_RECORD_VIDEO      ChatAction = "record_video"
	CHAT_ACTION_UPLOAD_VIDEO      ChatAction = "upload_video"
	CHAT_ACTION_RECORD_AUDIO      ChatAction = "record_audio"
	CHAT_ACTION_UPLOAD_AUDIO      ChatAction = "upload_audio"
	CHAT_ACTION_UPLOAD_DOCUMENT   ChatAction = "upload_document"
	CHAT_ACTION_FIND_LOCATION     ChatAction = "find_location"
	CHAT_ACTION_RECORD_VIDEO_NOTE ChatAction = "record_video_note"
	CHAT_ACTION_UPLOAD_VIDEO_NOTE ChatAction = "upload_video_note"

	MEMBER_STATUS_CREATOR       MemberStatus = "creator"
	MEMBER_STATUS_ADMINISTRATOR MemberStatus = "administrator"
	MEMBER_STATUS_MEMBER        MemberStatus = "member"
	MEMBER_STATUS_LEFT          MemberStatus = "left"
	MEMBER_STATUS_KICKED        MemberStatus = "kicked"

	MESSAGE_ENTITY_MENTION      MessageEntityType = "mention"
	MESSAGE_ENTITY_HASHTAG      MessageEntityType = "hashtag"
	MESSAGE_ENTITY_BOT_COMMAND  MessageEntityType = "bot_command"
	MESSAGE_ENTITY_URL          MessageEntityType = "url"
	MESSAGE_ENTITY_EMAIL        MessageEntityType = "email"
	MESSAGE_ENTITY_BOLD         MessageEntityType = "bold"
	MESSAGE_ENTITY_ITALIC       MessageEntityType = "italic"
	MESSAGE_ENTITY_CODE         MessageEntityType = "code"
	MESSAGE_ENTITY_PRE          MessageEntityType = "pre"
	MESSAGE_ENTITY_TEXT_LINK    MessageEntityType = "text_link"
	MESSAGE_ENTITY_TEXT_MENTION MessageEntityType = "text_mention"
)

type ParseMode string
type ChatType string
type ChatAction string
type MemberStatus string
type MessageEntityType string

// User object represents a Telegram user, bot
type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type ChatID string

func (chatID *ChatID) UnmarshalJSON(value []byte) error {
	*chatID = ChatID(value)
	return nil
}

// Chat object represents a chat.
type Chat struct {
	ID   ChatID   `json:"id"`
	Type ChatType `json:"type"`

	// Optional
	Title            string           `json:"title,omitempty"`
	Username         string           `json:"username,omitempty"`
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Photo            *ChatPhoto       `json:"photo,omitempty"`
	Description      string           `json:"description,omitempty"`
	InviteLink       string           `json:"invite_link,omitempty"`
	PinnedMessage    *Message         `json:"pinned_message,omitempty"`
	Permissions      *ChatPermissions `json:"permissions,omitempty"`
	StickerSetName   string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet bool             `json:"can_set_sticker_set,omitempty"`
}

// Message object represents a message.
type Message struct {
	MessageID int64  `json:"message_id"`
	From      User   `json:"from"`
	Date      uint64 `json:"date"`
	Chat      Chat   `json:"chat"`

	// Optional
	ForwardFrom           *User                `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat                `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  int64                `json:"forward_from_message_id,omitempty"`
	ForwardSignature      string               `json:"forward_signature,omitempty"`
	ForwardSenderName     string               `json:"forward_sender_name,omitempty"`
	ForwardDate           uint64               `json:"forward_date,omitempty"`
	ReplyToMessage        *Message             `json:"reply_to_message,omitempty"`
	EditDate              uint64               `json:"edit_date,omitempty"`
	MediaGroupID          string               `json:"media_group_id,omitempty"`
	AuthorSignature       string               `json:"author_signature,omitempty"`
	Text                  string               `json:"text,omitempty"`
	Entities              []MessageEntity      `json:"entities,omitempty"`
	CaptionEntities       []MessageEntity      `json:"caption_entities,omitempty"`
	Audio                 *Audio               `json:"audio,omitempty"`
	Document              *Document            `json:"document,omitempty"`
	Animation             *Animation           `json:"animation,omitempty"`
	Game                  *Game                `json:"game,omitempty"`
	Photo                 []PhotoSize          `json:"photo,omitempty"`
	Sticker               *Sticker             `json:"sticker,omitempty"`
	Video                 *Video               `json:"video,omitempty"`
	Voice                 *Voice               `json:"voice,omitempty"`
	VideoNote             *VideoNote           `json:"video_note,omitempty"`
	Caption               string               `json:"caption,omitempty"`
	Contact               *Contact             `json:"contact,omitempty"`
	Location              *Location            `json:"location,omitempty"`
	Venue                 *Venue               `json:"venue,omitempty"`
	Poll                  *Poll                `json:"poll,omitempty"`
	NewChatMembers        []User               `json:"new_chat_members,omitempty"`
	LeftChatMember        *User                `json:"left_chat_member,omitempty"`
	NewChatTitle          string               `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize          `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool                 `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool                 `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool                 `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool                 `json:"channel_chat_created,omitempty"`
	MigrateToChatID       ChatID               `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     ChatID               `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message             `json:"pinned_message,omitempty"`
	Invoice               *Invoice             `json:"invoice,omitempty"`
	SuccessfulPayment     *SuccessfulPayment   `json:"successful_payment,omitempty"`
	ConnectedWebsite      string               `json:"connected_website,omitempty"`
	PassportData          *PassportData        `json:"passport_data,omitempty"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// MessageEntity object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   MessageEntityType `json:"type"`
	Offset int               `json:"offset"`
	Limit  int               `json:"limit"`

	// Optional
	URL  string `json:"url,omitempty"`  // For “text_link” only, url that will be opened after user taps on the text
	User *User  `json:"user,omitempty"` // For “text_mention” only, the mentioned user
}

// PhotoSize object represents an image/sticker of a particular size.
type PhotoSize struct {
	FileID string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`

	// Optional
	FileSize uint64 `json:"file_size,omitempty"`
}

// Audio object represents an audio file (voice note).
type Audio struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`

	// Optional
	Performer string     `json:"performer,omitempty"`
	Title     string     `json:"title,omitempty"`
	MimeType  string     `json:"mime_type,omitempty"`
	Thumb     *PhotoSize `json:"thumb,omitempty"`
	FileSize  uint64     `json:"file_size,omitempty"`
}

// Document object represents a general file (as opposed to Photo or Audio).
// Telegram users can send files of any type of up to 1.5 GB in size.
type Document struct {
	FileID string `json:"file_id"`

	// Optional
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize uint64     `json:"file_size,omitempty"`
}

// Video object represents an MP4-encoded video.
type Video struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`

	// Optional
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize uint64     `json:"file_size,omitempty"`
}

// Animation object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`

	// Optional
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize *int       `json:"file_size,omitempty"`
}

// Voice object represents a voice note.
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`

	// Optional
	MimeType string `json:"mime_type,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
}

// VideoNote object represents a video message.
type VideoNote struct {
	FileID   string `json:"file_id"`
	Length   int    `json:"length"`
	Duration int    `json:"duration"`

	// Optional
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Contact object represents a contact to Telegram user
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`

	// Optional
	LastName string `json:"last_name,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
	VCard    string `json:"vcard,omitempty"`
}

// Location object represents geographic position.
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Venue object represents a venue.
type Venue struct {
	Location Location `json:"location"`
	Title    string   `json:"title"`
	Address  string   `json:"address"`

	// Optional
	FoursquareID   string `json:"foursquare_id,omitempty"`
	FoursquareType string `json:"foursquare_type,omitempty"`
}

// PollOption object contains information about one answer option in a poll.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// Poll object contains information about a poll.
type Poll struct {
	ID       string       `json:"id"`
	Question string       `json:"question"`
	Options  []PollOption `json:"options"`
	IsClosed bool         `json:"is_closed"`
}

// UserProfilePhotos object represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File object represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
type File struct {
	FileID string `json:"file_id"`

	// Optional
	FileSize uint64 `json:"file_size,omitempty"`
	FilePath string `json:"file_path,omitempty"`
}

type ReplyMarkup interface {
	itsReplyMarkup()
}

type replyMarkupImplementation struct{}

func (r replyMarkupImplementation) itsReplyMarkup() {}

// KeyboardButton object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields are mutually exclusive.
type KeyboardButton struct {
	Text string `json:"text"`

	// Optional
	RequestContact  bool `json:"request_contact,omitempty"`
	RequestLocation bool `json:"request_location,omitempty"`
}

// ReplyKeyboardMarkup object represents a custom keyboard with reply options
type ReplyKeyboardMarkup struct {
	replyMarkupImplementation
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
	Selective       bool               `json:"selective,omitempty"`
}

// ReplyKeyboardRemove object
// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button
type ReplyKeyboardRemove struct {
	replyMarkupImplementation
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"`
	Selective      bool `json:"selective,omitempty"`
}

// InlineKeyboardButton object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text string `json:"text,omitempty"`

	// Optional
	URL                          string    `json:"url,omitempty"`
	LoginURL                     *LoginURL `json:"login_url,omitempty"`
	CallbackData                 string    `json:"callback_data,omitempty"`
	SwitchInlineQuery            string    `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string    `json:"switch_inline_query_current_chat,omitempty"`
	Pay                          bool      `json:"pay,omitempty"`
}

// LoginURL object represents a parameter of the inline keyboard button used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in:
type LoginURL struct {
	URL string `json:"url"`

	// Optional
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// InlineKeyboardMarkup object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	replyMarkupImplementation
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// ChosenInlineResult object represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

// CallbackQuery object represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be presented.
// If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be presented.
type CallbackQuery struct {
	ID   string `json:"id"`
	From User   `json:"from"`

	// Optional
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance,omitempty"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// ForceReply object
// Upon receiving a message with this object,
// Telegram clients will display a reply interface to the user.
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	replyMarkupImplementation
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// ChatPhoto object represents a chat photo.
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

// ChatMember object contains information about one member of a chat.
type ChatMember struct {
	User   User         `json:"user"`
	Status MemberStatus `json:"status"`

	// Optional
	UntilDate             int64 `json:"until_date,omitempty"`
	CanBeEdited           bool  `json:"can_be_edited,omitempty"`
	CanPostMessages       bool  `json:"can_post_messages,omitempty"`
	CanEditMessages       bool  `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     bool  `json:"can_delete_messages,omitempty"`
	CanRestrictMembers    bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool  `json:"can_change_info,omitempty"`
	CanInviteUsers        bool  `json:"can_invite_users,omitempty"`
	CanPinMessages        bool  `json:"can_pin_messages,omitempty"`
	IsMember              bool  `json:"is_member,omitempty"`
	CanSendMessages       bool  `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool  `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool  `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool  `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool  `json:"can_add_web_page_previews,omitempty"`
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}

// ResponseParameters contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID ChatID `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int64  `json:"retry_after,omitempty"`
}

// Update object represents an incoming update.
// Only one of the optional parameters can be present in any given update.
type Update struct {
	UpdateID uint64 `json:"update_id"`

	// Optional
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
}

// WebhookInfo contains information about the current status of a webhook.
type WebhookInfo struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	LastErrorDate        uint64   `json:"last_error_date,omitempty"`
	LastErrorMessage     string   `json:"last_error_message,omitempty"`
	MaxConnections       int      `json:"max_connections,omitempty"`
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`
}
