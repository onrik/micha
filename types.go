package micha

const (
	PARSE_MODE_DEFAULT  ParseMode = ""
	PARSE_MODE_HTML     ParseMode = "HTML"
	PARSE_MODE_MARKDOWN ParseMode = "Markdown"

	CHAT_ACTION_TYPING          ChatAction = "typing"
	CHAT_ACTION_UPLOAD_PHOTO    ChatAction = "upload_photo"
	CHAT_ACTION_RECORD_VIDEO    ChatAction = "record_video"
	CHAT_ACTION_UPLOAD_VIDEO    ChatAction = "upload_video"
	CHAT_ACTION_RECORD_AUDIO    ChatAction = "record_audio"
	CHAT_ACTION_UPLOAD_AUDIO    ChatAction = "upload_audio"
	CHAT_ACTION_UPLOAD_DOCUMENT ChatAction = "upload_document"
	CHAT_ACTION_FIND_LOCATION   ChatAction = "find_location"
)

type ParseMode string
type ChatAction string

// User object represents a Telegram user, bot
type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// Chat object represents a Telegram user, bot or group chat.
type Chat struct {
	Id        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type FileBase struct {
	FileId   string `json:"file_id"`
	FileSize uint64 `json:"file_size"`
}

// Thumbnail object represents an image/sticker of a particular size.
type PhotoSize struct {
	FileBase

	Width  int `json:"width"`
	Height int `json:"height"`
}

// Photo object represents a photo with caption.
type Photo struct {
	FileBase
	PhotoSize

	Caption string `json:"caption"`
}

// Audio object represents an audio file (voice note).
type Audio struct {
	FileBase

	Duration  int    `json:"duration"`
	MimeType  string `json:"mime_type"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
}

// Document object represents a general file (as opposed to Photo or Audio).
// Telegram users can send files of any type of up to 1.5 GB in size.
type Document struct {
	FileBase

	Thumb    *PhotoSize `json:"thumb"`
	FileName string     `json:"file_name"`
	MimeType string     `json:"mime_type"`
}

// Sticker object represents a WebP image, so-called sticker.
type Sticker struct {
	FileBase

	Width  int        `json:"width"`
	Height int        `json:"height"`
	Thumb  *PhotoSize `json:"thumb"`
}

// Video object represents an MP4-encoded video.
type Video struct {
	FileBase

	Duration int        `json:"duration"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Thumb    *PhotoSize `json:"thumb"`
}

// Voice object represents a voice note.
type Voice struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

// Contact object represents a contact to Telegram user
type Contact struct {
	UserId      int64  `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

// Location object represents geographic position.
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Venue object represents a venue.
type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareId string   `json:"foursquare_id"`
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type File struct {
	FileBase
	FilePath string `json:"file_path"`
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Url    string `json:"url"`
}

// Message object represents a message.
type Message struct {
	MessageId int64  `json:"message_id"`
	From      User   `json:"from"`
	Date      uint64 `json:"date"`
	Chat      Chat   `json:"chat"`

	// Optional
	ForwardFrom           *User           `json:"forward_from"`
	ForwardDate           uint64          `json:"forward_date"`
	ReplyToMessage        *Message        `json:"reply_to_message"`
	Text                  string          `json:"text"`
	Entities              []MessageEntity `json:"entities"`
	Audio                 *Audio          `json:"audio"`
	Document              *Document       `json:"document"`
	Photo                 []PhotoSize     `json:"photo"`
	Sticker               *Sticker        `json:"sticker"`
	Video                 *Video          `json:"video"`
	Voice                 *Voice          `json:"voice"`
	Caption               string          `json:"caption"`
	Contact               *Contact        `json:"contact"`
	Location              *Location       `json:"location"`
	Venue                 *Venue          `json:"venue"`
	NewChatMember         *User           `json:"new_chat_member"`
	LeftChatMember        *User           `json:"left_chat_member"`
	NewChatTitle          string          `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo"`
	DeleteChatPhoto       bool            `json:"delete_chat_photo"`
	GroupChatCreated      bool            `json:"group_chat_created"`
	SupergroupChatCreated bool            `json:"supergroup_chat_created"`
	ChannelChatCreated    bool            `json:"channel_chat_created"`
	MigrateToChatId       int64           `json:"migrate_to_chat_id"`
	MigrateFromChatId     int64           `json:"migrate_from_chat_id"`
	PinnedMessage         *Message        `json:"pinned_message"`
}

type ReplyMarkup interface {
	_ItsReplyMarkup()
}

type ReplyMarkupImplementation struct{}

func (r ReplyMarkupImplementation) _ItsReplyMarkup() {}

// This object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields are mutually exclusive.
type KeyboardButton struct {
	Text string `json:"text"`

	// Optional
	RequestContact  bool `json:"request_contact,omitempty"`
	RequestLocation bool `json:"request_location,omitempty"`
}

// This object represents a custom keyboard with reply options
type ReplyKeyboardMarkup struct {
	ReplyMarkupImplementation
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
	Selective       bool               `json:"selective,omitempty"`
}

// Upon receiving a message with this object,
// Telegram clients will hide the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button
type ReplyKeyboardHide struct {
	ReplyMarkupImplementation
	HideKeyboard bool `json:"hide_keyboard,omitempty"`
	Selective    bool `json:"selective,omitempty"`
}

// This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text string `json:"text,omitempty"`

	// Optional
	Url               string `json:"url,omitempty"`
	CallbackData      string `json:"callback_data,omitempty"`
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
}

// This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	ReplyMarkupImplementation
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// Upon receiving a message with this object,
// Telegram clients will display a reply interface to the user.
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ReplyMarkupImplementation
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageId string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

// This object represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot, the field message will be presented.
// If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be presented.
type CallbackQuery struct {
	Id      string  `json:"id"`
	From    User    `json:"from"`
	Message Message `json:"message"`
	Data    string  `json:"data"`
}

// This object represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	Id     string `json:"id"`
	From   User   `json:"from"`
	Query  string `json:"query"`
	Offset string `json:"offset"`
}

// This object represents an incoming update.
// Only one of the optional parameters can be present in any given update.
type Update struct {
	UpdateId uint64 `json:"update_id"`

	// Optional
	Message            *Message            `json:"message"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
}
