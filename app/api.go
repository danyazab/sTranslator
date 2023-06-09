package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"translator/configs"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	messages chan tgbotapi.Update
	update   tgbotapi.Update
	updates  tgbotapi.UpdatesChannel
	config   *configs.Config

	speech          string
	id              int64
	callbackChannel chan string

	choosImput  chan string
	choosTransl chan string
	sourceLang  string
	targetLang  string
}

func NewBot(config *configs.Config) *Bot {
	return &Bot{config: config, choosImput: make(chan string), callbackChannel: make(chan string), choosTransl: make(chan string)}
}

func (s *Bot) Start() error {

	bot, err := tgbotapi.NewBotAPI(s.config.BotToken)
	if err != nil {
		return err
	}
	s.bot = bot

	s.GetUpdateTelegramBot()

	return nil
}
