package micha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// HttpClient interface
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HTTPError struct {
	StatusCode int
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("http status %d (%s)", e.StatusCode, http.StatusText(e.StatusCode))
}

type fileField struct {
	Source    io.Reader
	Fieldname string
	Filename  string
}

func handleResponse(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	if response.StatusCode > http.StatusBadRequest {
		return nil, HTTPError{response.StatusCode}
	}

	return io.ReadAll(response.Body)
}

func newGetRequest(ctx context.Context, url string, params url.Values) (*http.Request, error) {
	if params != nil {
		url += fmt.Sprintf("?%s", params.Encode())
	}
	return http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
}

func newPostRequest(ctx context.Context, url string, data interface{}) (*http.Request, error) {
	body := new(bytes.Buffer)
	if data != nil {
		if err := json.NewEncoder(body).Encode(data); err != nil {
			return nil, fmt.Errorf("encode data error: %w", err)
		}
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

func newMultipartRequest(url string, file *fileField, params url.Values) (*http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if file != nil {
		part, err := writer.CreateFormFile(file.Fieldname, file.Filename)
		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(part, file.Source); err != nil {
			return nil, err
		}
	}

	for field, values := range params {
		for i := range values {
			if err := writer.WriteField(field, values[i]); err != nil {
				return nil, err
			}
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())

	return request, nil
}
