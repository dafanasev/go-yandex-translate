package yandex_translate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	URL_ROOT       = "https://translate.yandex.net/api/v1.5/tr.json"
	LANGS_PATH     = "getLangs"
	TRANSLATE_PATH = "translate"
)

type Translator struct {
	apiKey string
}

type Languages struct {
	Code    int
	Message string
	Dirs    []string
	Langs   map[string]string
}

type Response struct {
	Code     int
	Message  string
	Lang     string
	Text     []string
	Detected map[string]string
}

func New(apiKey string) *Translator {
	return &Translator{apiKey: apiKey}
}

func (tr *Translator) GetLangs(ui string) (*Languages, error) {
	resp, err := http.PostForm(absUrl(LANGS_PATH), url.Values{"key": {tr.apiKey}, "ui": {ui}})
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

func (tr *Translator) Translate(lang, text string) (*Response, error) {
	builtParams := url.Values{"key": {tr.apiKey}, "lang": {lang}, "text": {text}, "options": {"1"}}
	resp, err := http.PostForm(absUrl(TRANSLATE_PATH), builtParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	if response.Code != 200 {
		return nil, fmt.Errorf("(%v) %v", response.Code, response.Message)
	}

	return &response, nil
}

func (response *Response) Result() string {
	return response.Text[0]
}

func absUrl(route string) string {
	return URL_ROOT + "/" + route
}
