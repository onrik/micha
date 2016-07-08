package http

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

type File struct {
	File      io.ReadCloser
	Fieldname string
	Filename  string
}

func handleResponse(response *http.Response) ([]byte, error) {
	if response.StatusCode > http.StatusBadRequest {
		return nil, fmt.Errorf("Response status: %d", response.StatusCode)
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		return handleResponse(response)
	}
}

func Post(url string, data interface{}) ([]byte, error) {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(data); err != nil {
		return nil, fmt.Errorf("Encode data error (%s)", err.Error())
	}

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return nil, err
	} else {
		return handleResponse(response)
	}
}

func PostMultipart(url string, file *File, params url.Values) ([]byte, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(file.Fieldname, file.Filename)
	if err != nil {
		return nil, err
	}

	defer file.File.Close()
	if _, err := io.Copy(part, file.File); err != nil {
		return nil, err
	}

	for field, values := range params {
		if len(values) > 0 {
			if err := writer.WriteField(field, values[0]); err != nil {
				return nil, err
			}
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return nil, err
	} else {
		return handleResponse(response)
	}
}
