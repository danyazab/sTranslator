package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	lang "translator/app/language"
	tr "translator/app/translate"
)

func (s *Bot) handleLang() error {
	msg := tgbotapi.NewMessage(s.id, lang.Langu)
	msg.ReplyMarkup = s.BoardLang()

	s.bot.Send(msg)

	return nil

}
func (s *Bot) hendleStart() error {
	msg := tgbotapi.NewMessage(s.id, tr.Translate(lang.En, s.speech, lang.RespWelcome1))
	s.bot.Send(msg)

	msg = tgbotapi.NewMessage(s.id, tr.Translate(lang.En, s.speech, lang.RespWelcome2))
	msg.ReplyMarkup = s.BoardMenu()

	s.bot.Send(msg)

	return nil

}
func (s *Bot) hendleImputS() error {
	go s.SendKeyboard(tr.Translate(lang.En, s.speech, lang.SelectImput), s.BoardLangList())
	go func() {
		s.choosImput <- "input"
	}()
	return nil

}
func (s *Bot) hendleTranslationS() error {
	go s.SendKeyboard(tr.Translate(lang.En, s.speech, lang.SelectTransla), s.BoardLangList())
	go func() {
		s.choosTransl <- "translation"
	}()

	return nil

}

func (s *Bot) hendleTranslate(message *tgbotapi.Message) error {
	go s.Send(tr.Translate(s.sourceLang, s.targetLang, message.Text))

	return nil

}
func (s *Bot) hendleAboutBot() error {
	go s.Send(tr.Translate(lang.En, s.speech, lang.RespAboutBot))

	return nil

}
