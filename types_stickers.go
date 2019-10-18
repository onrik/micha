package micha

// Sticker object represents a WebP image, so-called sticker.
type Sticker struct {
	FileID     string `json:"file_id"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	IsAnimated bool   `json:"is_animated,omitempty"`

	// Optional
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        string        `json:"emoji,omitempty"`
	SetName      string        `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     uint64        `json:"file_size,omitempty"`
}

// StickerSet object represents a sticker set.
type StickerSet struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	IsAnimated    bool      `json:"is_animated"`
	ContainsMasks bool      `json:"contains_masks"`
	Stickers      []Sticker `json:"stickers"`
}

// MaskPosition object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}
