package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type errorStruct struct{}

func (e errorStruct) MarshalJSON() ([]byte, error) {
	return nil, errors.New("marshal error")
}

type HttpTestSuite struct {
	suite.Suite
}

func (s *HttpTestSuite) SetupSuite() {
	httpmock.Activate()
}

func (s *HttpTestSuite) TearDownSuite() {
	httpmock.Deactivate()
}

func (s *HttpTestSuite) TearDownTest() {
	httpmock.Reset()
}

func (s *HttpTestSuite) TestHandleResponse() {
	data, err := handleResponse(httpmock.NewStringResponse(200, "ok"))
	s.Equal(nil, err)
	s.Equal([]byte("ok"), data)

	data, err = handleResponse(httpmock.NewStringResponse(400, "not ok"))
	s.Equal(nil, err)
	s.Equal([]byte("not ok"), data)

	data, err = handleResponse(httpmock.NewStringResponse(401, "not error"))
	s.Equal("Response status: 401", err.Error())
	s.Equal(0, len(data))
}

func (s *HttpTestSuite) TestSuccessGet() {
	url := "http://example.com"

	httpmock.RegisterResponder("GET", url, func(request *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(200, "111"), nil
	})

	response, err := Get(url)
	s.Equal(nil, err)
	s.Equal([]byte("111"), response)
}

func (s *HttpTestSuite) TestErrorGet() {
	url := "http://example.com"
	httpmock.RegisterResponder("GET", url, func(request *http.Request) (*http.Response, error) {
		return nil, errors.New("error")
	})

	response, err := Get(url)
	s.NotEqual(nil, err)
	s.Equal("Get http://example.com: error", err.Error())
	s.Equal(0, len(response))
}

func (s *HttpTestSuite) TestSuccessPost() {
	url := "http://example.com"
	data := map[string]string{
		"foo": "bar",
	}

	httpmock.RegisterResponder("POST", url, func(request *http.Request) (*http.Response, error) {
		requestData := map[string]string{}
		defer request.Body.Close()

		err := json.NewDecoder(request.Body).Decode(&requestData)
		s.Equal(nil, err)
		s.Equal(data, requestData)
		s.Equal("application/json", request.Header.Get("Content-Type"))
		return httpmock.NewStringResponse(200, "ok"), nil
	})

	response, err := Post(url, data)
	s.Equal(nil, err)
	s.Equal([]byte("ok"), response)
}

func (s *HttpTestSuite) TestErrorPost() {
	url := "http://example.com"
	httpmock.RegisterResponder("POST", url, func(request *http.Request) (*http.Response, error) {
		return nil, errors.New("error")
	})

	data := errorStruct{}
	response, err := Post(url, data)
	s.NotEqual(nil, err)
	s.True(strings.HasPrefix(err.Error(), "Encode data error"))

	response, err = Post(url, nil)
	s.NotEqual(nil, err)
	s.Equal("Post http://example.com: error", err.Error())
	s.Equal(0, len(response))
}

func (s *HttpTestSuite) TestSuccessPostMultipart() {
	data := url.Values{
		"foo": {"bar"},
	}
	url := "http://example.com"

	httpmock.RegisterResponder("POST", url, func(request *http.Request) (*http.Response, error) {
		contentType := request.Header.Get("Content-Type")
		s.True(strings.HasPrefix(contentType, "multipart/form-data; boundary="))

		defer request.Body.Close()
		err := request.ParseMultipartForm(1024)
		s.Equal(nil, err)
		s.Equal(request.MultipartForm.Value["foo"], []string{"bar"})

		files := request.MultipartForm.File["file"]
		s.Equal(1, len(files))
		s.Equal("somefile.ext", files[0].Filename)
		file, err := files[0].Open()
		s.Equal(nil, err)

		defer file.Close()
		data, err := ioutil.ReadAll(file)
		s.Equal(nil, err)
		s.Equal([]byte("filedata"), data)

		return httpmock.NewStringResponse(200, "ok"), nil
	})

	file := &File{
		Source:    bytes.NewBufferString("filedata"),
		Fieldname: "file",
		Filename:  "somefile.ext",
	}
	response, err := PostMultipart(url, file, data)
	s.Equal(nil, err)
	s.Equal([]byte("ok"), response)
}

func (s *HttpTestSuite) TestErrorPostMultipart() {
	url := "http://example.com"
	httpmock.RegisterResponder("POST", url, func(request *http.Request) (*http.Response, error) {
		return nil, errors.New("error")
	})

	file := &File{
		Source:    bytes.NewBufferString("filedata"),
		Fieldname: "file",
		Filename:  "somefile.ext",
	}
	response, err := PostMultipart(url, file, nil)
	s.NotEqual(nil, err)
	s.Equal("Post http://example.com: error", err.Error())
	s.Equal(0, len(response))
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
