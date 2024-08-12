package micha

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithAPIServer(t *testing.T) {
	options := &Options{}
	WithAPIServer("http://127.0.0.1/")(options)

	require.Equal(t, "http://127.0.0.1", options.apiServer)
}

func TestWithCtx(t *testing.T) {
	options := &Options{}
	ctx := context.WithValue(context.Background(), "foo", "bar")
	WithCtx(ctx)(options)

	require.Equal(t, ctx, options.ctx)
}
