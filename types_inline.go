package micha

const (
	INLINE_TYPE_RESULT_ARTICLE   InlineResultType = "article"
	INLINE_TYPE_RESULT_PHOTO     InlineResultType = "photo"
	INLINE_TYPE_RESULT_GIF       InlineResultType = "gif"
	INLINE_TYPE_RESULT_MPEG4_GIF InlineResultType = "mpeg4_gif"
	INLINE_TYPE_RESULT_VIDEO     InlineResultType = "video"
	INLINE_TYPE_RESULT_AUDIO     InlineResultType = "audio"
	INLINE_TYPE_RESULT_VOICE     InlineResultType = "voice"
	INLINE_TYPE_RESULT_DOCUMENT  InlineResultType = "document"
	INLINE_TYPE_RESULT_LOCATION  InlineResultType = "location"
	INLINE_TYPE_RESULT_VENUE     InlineResultType = "venue"
	INLINE_TYPE_RESULT_CONTACT   InlineResultType = "contact"
	INLINE_TYPE_RESULT_STICKER   InlineResultType = "sticker"
	INLINE_TYPE_RESULT_GAME      InlineResultType = "game"
)

// InlineQuery object represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

type InlineResultType string

type InlineQueryResults []InlineQueryResult

type InlineQueryResult interface {
	itsInlineQueryResult()
}

type inlineQueryResultImplementation struct{}

func (i inlineQueryResultImplementation) itsInlineQueryResult() {}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	inlineQueryResultImplementation
	Type  InlineResultType `json:"type"`
	ID    string           `json:"id"`
	Title string           `json:"title"`

	// Optional
	URL                 string                `json:"url,omitempty"`
	HideURL             bool                  `json:"hide_url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	inlineQueryResultImplementation
	Type     InlineResultType `json:"type"`
	ID       string           `json:"id"`
	PhotoURL string           `json:"photo_url"`

	// Optional
	MimeType            string                `json:"mime_type,omitempty"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	inlineQueryResultImplementation
	Type        InlineResultType `json:"type"`
	ID          string           `json:"id"`
	PhotoFileID string           `json:"photo_file_id"`

	// Optional
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file.
// By default, this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	inlineQueryResultImplementation
	Type   InlineResultType `json:"type"`
	ID     string           `json:"id"`
	GifURL string           `json:"gif_url"`

	// Optional
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	GifDuration         int                   `json:"gif_duration"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	inlineQueryResultImplementation
	Type      InlineResultType `json:"type"`
	ID        string           `json:"id"`
	GifFileID string           `json:"gif_file_id"`

	// Optional
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	inlineQueryResultImplementation
	Type     InlineResultType `json:"type"`
	ID       string           `json:"id"`
	Mpeg4URL string           `json:"mpeg4_url"`

	// Optional
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int                   `json:"mpeg4_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	inlineQueryResultImplementation
	Type        InlineResultType `json:"type"`
	ID          string           `json:"id"`
	Mpeg4FileID string           `json:"mpeg4_file_id"`

	// Optional
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	inlineQueryResultImplementation
	Type     InlineResultType `json:"type"`
	ID       string           `json:"id"`
	VideoURL string           `json:"video_url"`
	MimeType string           `json:"mime_type"`

	// Optional
	ThumbURL            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	VideoWidth          int                   `json:"video_width,omitempty"`
	VideoHeight         int                   `json:"video_height,omitempty"`
	VideoDuration       int                   `json:"video_duration,omitempty"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	inlineQueryResultImplementation
	Type        InlineResultType `json:"type"`
	ID          string           `json:"id"`
	VideoFileID string           `json:"video_file_id"`

	// Optional
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to an mp3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	inlineQueryResultImplementation
	Type     InlineResultType `json:"type"`
	ID       string           `json:"id"`
	AudioURL string           `json:"audio_url"`
	Title    string           `json:"title"`

	// Optional
	Performer           string                `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to an mp3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	inlineQueryResultImplementation
	Type        InlineResultType `json:"type"`
	ID          string           `json:"id"`
	AudioFileID string           `json:"audio_file_id"`

	// Optional
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a voice recording in an .ogg container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	inlineQueryResultImplementation
	Type     InlineResultType `json:"type"`
	ID       string           `json:"id"`
	VoiceURL string           `json:"voice_url"`
	Title    string           `json:"title"`

	// Optional
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	inlineQueryResultImplementation
	Type        InlineResultType `json:"type"`
	ID          string           `json:"id"`
	VoiceFileID string           `json:"voice_file_id"`

	// Optional
	Title               string                `json:"title"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	inlineQueryResultImplementation
	Type        InlineResultType `json:"type"`
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	DocumentURL string           `json:"document_url"`
	MimeType    string           `json:"mime_type"`

	// Optional
	Caption             string                `json:"caption,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Currently, only pdf-files and zip archives can be sent using this method.
type InlineQueryResultCachedDocument struct {
	inlineQueryResultImplementation
	Type           InlineResultType `json:"type"`
	ID             string           `json:"id"`
	Title          string           `json:"title"`
	DocumentFileID string           `json:"document_file_id"`

	// Optional
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a location on a map.
// By default, the location will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	inlineQueryResultImplementation
	Type      InlineResultType `json:"type"`
	ID        string           `json:"id"`
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
	Title     string           `json:"title"`

	// Optional
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a venue.
// By default, the venue will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	inlineQueryResultImplementation
	Type      InlineResultType `json:"type"`
	ID        string           `json:"id"`
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
	Title     string           `json:"title"`
	Address   string           `json:"address"`

	// Optional
	FoursquareID        string                `json:"foursquare_id,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	inlineQueryResultImplementation
	Type          InlineResultType `json:"type"`
	ID            string           `json:"id"`
	StickerFileID string           `json:"sticker_file_id"`

	// Optional
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a contact with a phone number.
// By default, this contact will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	inlineQueryResultImplementation
	Type        InlineResultType `json:"type"`
	ID          string           `json:"id"`
	PhoneNumber string           `json:"phone_number"`
	FirstName   string           `json:"first_name"`

	// Optional
	LastName            string                `json:"last_name,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a Game.
type InlineQueryResultGame struct {
	inlineQueryResultImplementation
	Type          InlineResultType `json:"type"`
	ID            string           `json:"id"`
	GameShortName string           `json:"game_short_name"`

	// Optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InputMessageContent interface {
	itsInputMessageContent()
}

type inputMessageContentImplementation struct{}

func (i inputMessageContentImplementation) itsInputMessageContent() {}

// InputTextMessageContent contains text for displaying as an inline query result.
type InputTextMessageContent struct {
	inputMessageContentImplementation
	MessageText           string    `json:"message_text"`
	ParseMode             ParseMode `json:"parse_mode"`
	DisableWebPagePreview bool      `json:"disable_web_page_preview"`
}

// InputLocationMessageContent contains a location for displaying as an inline query result.
type InputLocationMessageContent struct {
	inputMessageContentImplementation
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// InputVenueMessageContent contains a venue for displaying an inline query result.
type InputVenueMessageContent struct {
	inputMessageContentImplementation
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareID string  `json:"foursquare_id"`
}

// InputContactMessageContent contains a contact for displaying as an inline query result.
type InputContactMessageContent struct {
	inputMessageContentImplementation
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}
