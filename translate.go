// Package translate is the Yandex.Translate API client
package translate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	urlRoot       = "https://translate.yandex.net/api/v1.5/tr.json"
	langsPath     = "getLangs"
	translatePath = "translate"
)

// Translator holds api key
type Translator struct {
	apiKey string
}

// Languages holds GetLangs method response
type Languages struct {
	Code    int
	Message string
	Dirs    []string
	Langs   map[string]string
}

// Response holds Translate method response
type Response struct {
	Code     int
	Message  string
	Lang     string
	Text     []string
	Detected map[string]string
}

// New returns translator instance
func New(apiKey string) *Translator {
	return &Translator{apiKey: apiKey}
}

// GetLangs returns supported languages
func (tr *Translator) GetLangs(ui string) (*Languages, error) {
	resp, err := http.PostForm(absURL(langsPath), url.Values{"key": {tr.apiKey}, "ui": {ui}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Languages
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, fmt.Errorf("(%v) %v", response.Code, response.Message)
	}

	return &response, nil
}

// Translate returns translation for the request
func (tr *Translator) Translate(lang, text string) (*Response, error) {
	errMsg := fmt.Sprintf("can't get translation for %s", text)

	builtParams := url.Values{"key": {tr.apiKey}, "lang": {lang}, "text": {text}, "options": {"1"}}
	resp, err := http.PostForm(absURL(translatePath), builtParams)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	defer resp.Body.Close()

	var response Response
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	if response.Code != 200 {
		return nil, errors.Errorf("%s: %d, %s", errMsg, response.Code, response.Message)
	}

	return &response, nil
}

// Result returns translation as a string
func (response *Response) Result() string {
	return response.Text[0]
}

func absURL(route string) string {
	return urlRoot + "/" + route
}
