package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
	lang "translator/app/language"
)

func (s *Bot) handleLang() error {
	msg := tgbotapi.NewMessage(s.id, lang.Langu)
	msg.ReplyMarkup = s.BoardLang()
	s.speech = lang.En

	s.bot.Send(msg)

	return nil
}

func (s *Bot) handleStop() error {
	go s.Send(s.tr.Translate(lang.En, s.speech, lang.RespStop))
	return nil
}

func (s *Bot) hendleStart() error {
	s.bot.Send(tgbotapi.NewChatAction(s.id, tgbotapi.ChatTyping))
	time.Sleep(500 * time.Millisecond)
	msg := tgbotapi.NewMessage(s.id, s.tr.Translate(lang.En, s.speech, lang.RespWelcome1))
	s.bot.Send(msg)

	s.bot.Send(tgbotapi.NewChatAction(s.id, tgbotapi.ChatTyping))
	time.Sleep(350 * time.Millisecond)
	msg = tgbotapi.NewMessage(s.id, s.tr.Translate(lang.En, s.speech, lang.RespWelcome2))
	msg.ReplyMarkup = s.BoardMenu(s.speech, "en")

	s.bot.Send(msg)

	return nil

}
func (s *Bot) hendleImputS() error {
	go s.SendKeyboard(s.tr.Translate(lang.En, s.speech, lang.SelectImput), s.BoardLangList())
	go func() {
		s.choosImput <- "input"
	}()
	return nil

}
func (s *Bot) hendleTranslationS() error {
	go s.SendKeyboard(s.tr.Translate(lang.En, s.speech, lang.SelectTransla), s.BoardLangList())
	go func() {
		s.choosTransl <- "translation"
	}()

	return nil

}

func (s *Bot) hendleTranslate() error {
	go s.Send(s.tr.Translate(s.sourceLang, s.targetLang, s.update.Message.Text))

	return nil

}
func (s *Bot) hendleAboutBot() error {
	go s.Send(s.tr.Translate(lang.En, s.speech, lang.RespAboutBot))

	return nil

}
