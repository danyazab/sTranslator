package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	lang "translator/app/language"
	tr "translator/app/translate"
)

func (s *Bot) BoardLang() tgbotapi.ReplyKeyboardMarkup {

	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("uk"),
			tgbotapi.NewKeyboardButton("en"),
			tgbotapi.NewKeyboardButton("de"),
			tgbotapi.NewKeyboardButton("pl"),
			tgbotapi.NewKeyboardButton("es"),
			tgbotapi.NewKeyboardButton("fr"),
		),
	)
	return replyKeyboard

}
func (s *Bot) BoardMenu() tgbotapi.ReplyKeyboardMarkup {

	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(tr.Translate("en", s.speech, lang.SelectImput)),
			tgbotapi.NewKeyboardButton(tr.Translate("en", s.speech, lang.SelectTransla)),
		), tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(tr.Translate("en", s.speech, lang.AboutBot)),
		),
	)
	return replyKeyboard

}

func (s *Bot) BoardLangList() tgbotapi.InlineKeyboardMarkup {

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lang.Uk, lang.Uk),
			tgbotapi.NewInlineKeyboardButtonData(lang.En, lang.En),
			tgbotapi.NewInlineKeyboardButtonData(lang.De, lang.De),
			tgbotapi.NewInlineKeyboardButtonData(lang.Cs, lang.Cs),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lang.Pl, lang.Pl),
			tgbotapi.NewInlineKeyboardButtonData(lang.Es, lang.Es),
			tgbotapi.NewInlineKeyboardButtonData(lang.Fr, lang.Fr),
			tgbotapi.NewInlineKeyboardButtonData(lang.It, lang.It),
		),
	)
	return numericKeyboard

}
