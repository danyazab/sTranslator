package translate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
)

func encodeURI(s string) string {
	return url.QueryEscape(s)
}

func getTranslationURL(source, sourceLang, targetLang string) string {
	encodedSource := encodeURI(source)
	return fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s",
		sourceLang, targetLang, encodedSource)
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
	defer func() {
		if r := recover(); r != nil {
			// Обробка помилки
			fmt.Println("Помилка:", r)
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
		cText := strings.Join(text, "")

		// Convert the first character to uppercase
		//if len(cText) > 0 {
		//	cText = strings.ToUpper(string(cText[0])) + cText[1:]
		//}

		return cText
	} else {
		return "No translated data in response"
	}
}

func Translate(sourceLang, targetLang, source string) string {
	// Перевірка, чи рядок закодований у форматі UTF-8
	if !utf8.ValidString(source) {
		return "Source string is not encoded in UTF-8"
	}

	translationURL := getTranslationURL(source, sourceLang, targetLang)
	response, err := fetchTranslation(translationURL)
	if err != nil {
		return "err fetchTranslation"
	}

	result, err := parseTranslationResponse(response)
	if err != nil {
		return "err parseTranslationResponse"
	}

	return extractTranslatedText(result)
}
