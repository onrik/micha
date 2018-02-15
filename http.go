package micha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

// HttpClient interface
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type fileField struct {
	Source    io.Reader
	Fieldname string
	Filename  string
}

func handleResponse(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	if response.StatusCode > http.StatusBadRequest {
		return nil, fmt.Errorf("HTTP status: %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

func newGetRequest(url string, params url.Values) (*http.Request, error) {
	if params != nil {
		url += fmt.Sprintf("?%s", params.Encode())
	}
	return http.NewRequest(http.MethodGet, url, nil)
}

func newPostRequest(url string, data interface{}) (*http.Request, error) {
	body := new(bytes.Buffer)
	if data != nil {
		if err := json.NewEncoder(body).Encode(data); err != nil {
			return nil, fmt.Errorf("Encode data error (%s)", err.Error())
		}
	}

	request, err := http.NewRequest(http.MethodPost, url, body)
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
