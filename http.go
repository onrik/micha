package micha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode > http.StatusBadRequest {
		return nil, fmt.Errorf("Response status: %d", response.StatusCode)
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func post(url string, data interface{}) ([]byte, error) {
	buff := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buff).Encode(data); err != nil {
		return nil, fmt.Errorf("Encode data error (%s)", err.Error())
	}

	request, err := http.NewRequest("POST", url, buff)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode > http.StatusBadRequest {
		return nil, fmt.Errorf("Response status: %d", response.StatusCode)
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
