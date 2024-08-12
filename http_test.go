package micha

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPError(t *testing.T) {
	var err error = HTTPError{
		StatusCode: http.StatusNotFound,
	}

	require.Equal(t, "http status 404 (Not Found)", err.Error())
}

func TestHandleResponse(t *testing.T) {
	response := &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(nil)),
		StatusCode: http.StatusForbidden,
	}

	body, err := handleResponse(response)
	require.NotNil(t, err)
	require.Nil(t, body)

	require.Equal(t, "http status 403 (Forbidden)", err.Error())
}
