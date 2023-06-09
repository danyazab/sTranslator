package app

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	Registration = "Registration"

	startReply = "Привет, чтобы сохранить ссылку на своем Pocket аккаунте, для начала тебе неодходимо дать мне на это доступ. Для этого переходи по ссылке: \n%s"
)

func (s *Bot) GetUpdateTelegramBot() {

	s.bot.Debug = true
	log.Printf("Authorized on account %s", s.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := s.bot.GetUpdatesChan(u)

	for update := range updates {
		fmt.Println("d")

		if update.CallbackQuery != nil {
			s.callbackChannel <- "delete"
			go s.handleCommandCallbackQuery(update)
		}

		if update.Message == nil {

			continue
		}
		s.update = update
		s.updates = updates
		s.id = update.Message.Chat.ID
		if update.Message.IsCommand() {
			s.handleCommand(update.Message)
			continue
		}

		s.handleCommandText(update.Message)

		//go s.ControlerMessage(s.bot, updates, s.update)
	}
}

func (s *Bot) Send(text string) {
	msg := tgbotapi.NewMessage(s.id, text)
	msg.ParseMode = tgbotapi.ModeMarkdown

	s.bot.Send(tgbotapi.NewChatAction(s.id, tgbotapi.ChatTyping))
	time.Sleep(1 * time.Second)
	sentMsg, _ := s.bot.Send(msg)
	time.Sleep(35 * time.Second)

	delMsg := tgbotapi.NewDeleteMessage(s.id, sentMsg.MessageID)
	s.bot.DeleteMessage(delMsg)
}

func (s *Bot) SendKeyboard(text string, button tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(s.id, text)
	msg.ReplyMarkup = button

	msg.ParseMode = tgbotapi.ModeMarkdown

	s.bot.Send(tgbotapi.NewChatAction(s.id, tgbotapi.ChatTyping))
	time.Sleep(1 * time.Second)
	sentMsg, _ := s.bot.Send(msg)

	for {
		select {
		case <-time.After(30 * time.Second):
			s.deleteMessage(sentMsg.MessageID)
			return
		case <-s.callbackChannel:
			s.deleteMessage(sentMsg.MessageID)
			return
		}
	}
}

func (s *Bot) deleteMessage(messageID int) {
	time.Sleep(500 * time.Millisecond)
	delMsg := tgbotapi.NewDeleteMessage(s.id, messageID)
	s.bot.DeleteMessage(delMsg)
}
