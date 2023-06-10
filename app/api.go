package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"translator/app/translate"
	"translator/configs"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	update tgbotapi.Update
	config *configs.Config
	tr     *translate.Trsl

	speech          string
	id              int64
	callbackChannel chan string
	deleteChannel   chan string
	sentMsgID       int

	choosImput  chan string
	choosTransl chan string
	sourceLang  string
	targetLang  string
}

func NewBot(config *configs.Config) *Bot {
	trsl := translate.NewTrsl(config)

	return &Bot{config: config,
		choosImput:      make(chan string),
		callbackChannel: make(chan string),
		deleteChannel:   make(chan string),
		choosTransl:     make(chan string),
		tr:              trsl,
	}
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
