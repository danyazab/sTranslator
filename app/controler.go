package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	lang "translator/app/language"
)

func (s *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case lang.Start:
		return s.handleLang()
	case lang.Stop:
		return s.handleLang()
	case lang.SImput:
		return s.hendleImputS()
	case lang.STranslate:
		return s.hendleTranslationS() //
	case lang.About:
		return s.hendleAboutBot()
	default:
		return s.handleCommand(message)
	}
}

func (s *Bot) handleCommandText(message *tgbotapi.Message) error {
	if strings.Contains(message.Text, "⇆") {
		s.sourceLang, s.targetLang = s.targetLang, s.sourceLang
		s.SendRespChange()

	} else {
		switch message.Text {
		case lang.De, lang.Es, lang.En, lang.Fr, lang.Uk, lang.Pl:
			s.speech = message.Text
			s.sourceLang = message.Text
			s.targetLang = lang.En
			s.hendleStart()
			return s.SendRespChange()

		case s.tr.Translate(lang.En, s.speech, lang.SelectImput):

			return s.hendleImputS()
		case s.tr.Translate(lang.En, s.speech, lang.SelectTransla):
			return s.hendleTranslationS()

		case s.tr.Translate(lang.En, s.speech, lang.AboutBot):
			return s.hendleAboutBot()

		default:
			s.hendleTranslate()
		}
	}
	return nil
}

func (s *Bot) handleCommandCallbackQuery(u tgbotapi.Update) error {
	select {
	case <-s.choosTransl:
		s.targetLang = u.CallbackQuery.Data
		go s.Send(fmt.Sprintf("%s %s", s.tr.Translate("en", s.speech, lang.RespChoos), u.CallbackQuery.Data))
		s.SendRespChange()
	case <-s.choosImput:
		go s.Send(fmt.Sprintf("%s %s", s.tr.Translate("en", s.speech, lang.RespChoos), u.CallbackQuery.Data))
		s.SendRespChange()
		s.sourceLang = u.CallbackQuery.Data
	}
	return nil
}
