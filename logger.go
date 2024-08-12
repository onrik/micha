package micha

import (
	"context"
)

// Logger interface
type Logger interface {
	ErrorContext(ctx context.Context, msg string, args ...any)
}
