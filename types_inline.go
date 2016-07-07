package micha

const (
	INLINE_TYPE_ARTICLE  = "article"
	INLINE_TYPE_PHOTO    = "photo"
	INLINE_TYPE_GIF      = "gif"
	INLINE_TYPE_VIDEO    = "video"
	INLINE_TYPE_AUDIO    = "audio"
	INLINE_TYPE_DOCUMENT = "document"
	INLINE_TYPE_VOICE    = "voice"
	INLINE_TYPE_LOCATION = "location"
	INLINE_TYPE_STICKER  = "sticker"
)

type InlineQueryResults []InlineQueryResult

type InlineQueryResult interface {
	_ItsInlineQueryResult()
}

type InlineQueryResultImplementation struct{}

func (i InlineQueryResultImplementation) _ItsInlineQueryResult() {}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	InlineQueryResultImplementation
	Type  string `json:"type"`
	Id    string `json:"id"`
	Title string `json:"title"`

	// Optional
	Url                 string                `json:"url,omitempty"`
	HideUrl             bool                  `json:"hide_url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	InlineQueryResultImplementation
	Type     string `json:"type"`
	Id       string `json:"id"`
	PhotoUrl string `json:"photo_url"`

	// Optional
	MimeType            string                `json:"mime_type,omitempty"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
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
	InlineQueryResultImplementation
	Type        string `json:"type"`
	Id          string `json:"id"`
	PhotoFileId string `json:"photo_file_id"`

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
	InlineQueryResultImplementation
	Type   string `json:"type"`
	Id     string `json:"id"`
	GifUrl string `json:"gif_url"`

	// Optional
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	InlineQueryResultImplementation
	Type      string `json:"type"`
	Id        string `json:"id"`
	GifFileId string `json:"gif_file_id"`

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
	InlineQueryResultImplementation
	Type     string `json:"type"`
	Id       string `json:"id"`
	Mpeg4Url string `json:"mpeg4_url"`

	// Optional
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	InlineQueryResultImplementation
	Type        string `json:"type"`
	Id          string `json:"id"`
	Mpeg4FileId string `json:"mpeg4_file_id"`

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
	InlineQueryResultImplementation
	Type     string `json:"type"`
	Id       string `json:"id"`
	VideoUrl string `json:"video_url"`
	MimeType string `json:"mime_type"`

	// Optional
	ThumbUrl            string                `json:"thumb_url,omitempty"`
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
	InlineQueryResultImplementation
	Type        string `json:"type"`
	Id          string `json:"id"`
	VideoFileId string `json:"video_file_id"`

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
	InlineQueryResultImplementation
	Type     string `json:"type"`
	Id       string `json:"id"`
	AudioUrl string `json:"audio_url"`
	Title    string `json:"title"`

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
	InlineQueryResultImplementation
	Type        string `json:"type"`
	Id          string `json:"id"`
	AudioFileId string `json:"audio_file_id"`

	// Optional
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a voice recording in an .ogg container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	InlineQueryResultImplementation
	Type     string `json:"type"`
	Id       string `json:"id"`
	VoiceUrl string `json:"voice_url"`
	Title    string `json:"title"`

	// Optional
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	InlineQueryResultImplementation
	Type        string `json:"type"`
	Id          string `json:"id"`
	VoiceFileId string `json:"voice_file_id"`

	// Optional
	Title               string                `json:"title"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	InlineQueryResultImplementation
	Type        string `json:"type"`
	Id          string `json:"id"`
	Title       string `json:"title"`
	DocumentUrl string `json:"document_url"`
	MimeType    string `json:"mime_type"`

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
	InlineQueryResultImplementation
	Type           string `json:"type"`
	Id             string `json:"id"`
	Title          string `json:"title"`
	DocumentFileId string `json:"document_file_id"`

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
	InlineQueryResultImplementation
	Type      string  `json:"type"`
	Id        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`

	// Optional
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// Represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	InlineQueryResultImplementation
	Type          string `json:"type"`
	Id            string `json:"id"`
	StickerFileId string `json:"sticker_file_id"`

	// Optional
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

type InputMessageContent interface {
	_ItsInputMessageContent()
}

type InputMessageContentImplementation struct{}

func (i InputMessageContentImplementation) _ItsInputMessageContent() {}

// InputTextMessageContent contains text for displaying as an inline query result.
type InputTextMessageContent struct {
	InputMessageContentImplementation
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

// InputLocationMessageContent contains a location for displaying as an inline query result.
type InputLocationMessageContent struct {
	InputMessageContentImplementation
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// InputVenueMessageContent contains a venue for displaying an inline query result.
type InputVenueMessageContent struct {
	InputMessageContentImplementation
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareId string  `json:"foursquare_id"`
}

// InputContactMessageContent contains a contact for displaying as an inline query result.
type InputContactMessageContent struct {
	InputMessageContentImplementation
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}
