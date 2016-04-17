package micha

type InlineQueryResults []InlineQueryResult

type InlineQueryResult interface {
	_ItsInlineQueryResult()
}

type InlineQueryResultBase struct {
	Type string `json:"type"`
	Id   string `json:"id"`

	// Optional
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultBase) _ItsInlineQueryResult() {}

// InlineQueryResultArticle is an inline query response article.
type InlineQueryResultArticle struct {
	InlineQueryResultBase
	Title string `json:"title"`

	// Optional
	Url         string `json:"url"`
	HideUrl     bool   `json:"hide_url"`
	Description string `json:"description"`
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

// InlineQueryResultPhoto is an inline query response photo.
type InlineQueryResultPhoto struct {
	InlineQueryResultBase
	PhotoUrl string `json:"photo_url"`

	// Optional
	MimeType    string `json:"mime_type"`
	PhotoWidth  int    `json:"photo_width"`
	PhotoHeight int    `json:"photo_height"`
	ThumbUrl    string `json:"thumb_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Caption     string `json:"caption"`
}

// InlineQueryResultGIF is an inline query response GIF.
type InlineQueryResultGIF struct {
	InlineQueryResultBase
	GifUrl string `json:"gif_url"`

	// Optional
	GifWidth  int    `json:"gif_width"`
	GifHeight int    `json:"gif_height"`
	ThumbUrl  string `json:"thumb_url"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
}

// InlineQueryResultMPEG4GIF is an inline query response MPEG4 GIF.
type InlineQueryResultMPEG4GIF struct {
	InlineQueryResultBase
	Mpeg4Url string `json:"mpeg4_url"`

	// Optional
	Mpeg4Width  int    `json:"mpeg4_width"`
	Mpeg4Height int    `json:"mpeg4_height"`
	ThumbURL    string `json:"thumb_url"`
	Title       string `json:"title"`
	Caption     string `json:"caption"`
}

// InlineQueryResultVideo is an inline query response video.
type InlineQueryResultVideo struct {
	InlineQueryResultBase
	VideoUrl string `json:"video_url"`
	MimeType string `json:"mime_type"`

	// Optional
	ThumbUrl      string `json:"thumb_url"`
	Title         string `json:"title"`
	Caption       string `json:"caption"`
	VideoWidth    int    `json:"video_width"`
	VideoHeight   int    `json:"video_height"`
	VideoDuration int    `json:"video_duration"`
	Description   string `json:"description"`
}

// InlineQueryResultAudio is an inline query response audio.
type InlineQueryResultAudio struct {
	InlineQueryResultBase
	AudioUrl string `json:"audio_url"`
	Title    string `json:"title"`

	// Optional
	Performer     string `json:"performer"`
	AudioDuration int    `json:"audio_duration"`
}

// InlineQueryResultVoice is an inline query response voice.
type InlineQueryResultVoice struct {
	InlineQueryResultBase
	VoiceUrl string `json:"voice_url"`
	Title    string `json:"title"`

	// Optional
	VoiceDuration int `json:"voice_duration"`
}

// InlineQueryResultDocument is an inline query response document.
type InlineQueryResultDocument struct {
	InlineQueryResultBase
	Title       string `json:"title"`
	DocumentUrl string `json:"document_url"`
	MimeType    string `json:"mime_type"`

	// Optional
	Caption     string `json:"caption"`
	Description string `json:"description"`
	ThumbURL    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

// InlineQueryResultLocation is an inline query response location.
type InlineQueryResultLocation struct {
	InlineQueryResultBase
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`

	// Optional
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

type InputMessageContent interface {
	_ItsInputMessageContent()
}

type InputMessageContentBase struct{}

func (i InlineQueryResultBase) _ItsInputMessageContent() {}

// InputTextMessageContent contains text for displaying as an inline query result.
type InputTextMessageContent struct {
	InputMessageContentBase
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

// InputLocationMessageContent contains a location for displaying as an inline query result.
type InputLocationMessageContent struct {
	InputMessageContentBase
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// InputVenueMessageContent contains a venue for displaying an inline query result.
type InputVenueMessageContent struct {
	InputMessageContentBase
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareId string  `json:"foursquare_id"`
}

// InputContactMessageContent contains a contact for displaying as an inline query result.
type InputContactMessageContent struct {
	InputMessageContentBase
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}
