package micha

// Game object represents a game.
// Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Photo       []PhotoSize `json:"photo"`

	// Optional
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    *Animation      `json:"animation"`
}

// GameHighScore object represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}
