package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	lang "translator/app/language"
	tr "translator/app/translate"
)

func (s *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case lang.Start:
		return s.handleLang()
	case lang.Stop:
		return nil ///
	default:
		return s.handleCommand(message)
	}
}

func (s *Bot) handleCommandText(message *tgbotapi.Message) error {

	switch message.Text {
	case lang.De, lang.Es, lang.En, lang.Fr, lang.Uk, lang.Pl:
		s.speech = message.Text
		s.sourceLang = message.Text
		s.targetLang = "en"
		return s.hendleStart()

	case tr.Translate(lang.En, s.speech, lang.SelectImput):

		return s.hendleImputS()
	case tr.Translate(lang.En, s.speech, lang.SelectTransla):
		return s.hendleTranslationS()

	case tr.Translate(lang.En, s.speech, lang.AboutBot):
		return s.hendleAboutBot()

	default:
		s.hendleTranslate(message)
	}
	return nil
}

func (s *Bot) handleCommandCallbackQuery(u tgbotapi.Update) error {
	select {
	case <-s.choosTransl:
		s.targetLang = u.CallbackQuery.Data
		go s.Send(fmt.Sprintf("%s %s", tr.Translate("en", s.speech, lang.RespChoos), u.CallbackQuery.Data))
	case <-s.choosImput:
		go s.Send(fmt.Sprintf("%s %s", tr.Translate("en", s.speech, lang.RespChoos), u.CallbackQuery.Data))
		s.sourceLang = u.CallbackQuery.Data
	}
	return nil
}
