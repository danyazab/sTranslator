package translate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"translator/configs"
	"unicode/utf8"
)

type Trsl struct {
	config *configs.Config
}

func NewTrsl(config *configs.Config) *Trsl {
	return &Trsl{config: config}

}

func encodeURI(s string) string {
	return url.QueryEscape(s)
}

func (s *Trsl) getTranslationURL(source, sourceLang, targetLang string) string {
	encodedSource := encodeURI(source)
	return fmt.Sprintf("%s%s&tl=%s&dt=t&q=%s",
		s.config.Url, sourceLang, targetLang, encodedSource)
}

func fetchTranslation(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Error getting translation")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("Error reading response body")
	}

	return body, nil
}

func parseTranslationResponse(response []byte) ([]interface{}, error) {
	var result []interface{}
	err := json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func extractTranslatedText(result []interface{}) string {
	var cText string
	defer func() {
		if r := recover(); r != nil {
			cText = "translation error"
		}
	}()

	var text []string

	if len(result) > 0 {
		inner := result[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		cText = strings.Join(text, "")

		return cText
	} else {
		return "No translated data in response"
	}
}

func (s *Trsl) Translate(sourceLang, targetLang, source string) string {
	if !utf8.ValidString(source) {
		return "Source string is not encoded in UTF-8"
	}

	translationURL := s.getTranslationURL(source, sourceLang, targetLang)
	response, err := fetchTranslation(translationURL)
	if err != nil {
		return "err fetch Translation"
	}

	result, err := parseTranslationResponse(response)
	if err != nil {
		return "please restart bot"
	}

	return extractTranslatedText(result)
}
