package micha

import "fmt"

type Error struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d (%s)", e.Code, e.Description)
}
