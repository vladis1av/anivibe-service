package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func InitBot(botToken string, debug bool) error {
	var err error
	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return fmt.Errorf("failed to create Telegram bot: %w", err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	go startBot()

	return nil
}

func GetBot() *tgbotapi.BotAPI {
	return bot
}

func startBot() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "id":
				chatID := update.Message.Chat.ID
				msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Привет! Твой ChatID: %d", chatID))
				bot.Send(msg)
			}
		}
	}
}
