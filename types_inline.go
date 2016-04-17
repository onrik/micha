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
)

type InlineQueryResults []InlineQueryResult

type InlineQueryResult interface {
	_ItsInlineQueryResult()
}

type InlineQueryResultImplementation struct{}

func (i InlineQueryResultImplementation) _ItsInlineQueryResult() {}

// InlineQueryResultArticle is an inline query response article.
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

// InlineQueryResultPhoto is an inline query response photo.
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

// InlineQueryResultGIF is an inline query response GIF.
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

// InlineQueryResultMPEG4GIF is an inline query response MPEG4 GIF.
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

// InlineQueryResultVideo is an inline query response video.
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

// InlineQueryResultAudio is an inline query response audio.
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

// InlineQueryResultVoice is an inline query response voice.
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

// InlineQueryResultDocument is an inline query response document.
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

// InlineQueryResultLocation is an inline query response location.
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
