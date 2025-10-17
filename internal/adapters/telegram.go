package adapters

import (
	"context"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAdapter struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func NewTelegramAdapter(token string, chatID int64) (*TelegramAdapter, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.New("ошибка инициализации telegram-бота")
	}

	return &TelegramAdapter{bot: bot, chatID: chatID}, nil
}

func (t *TelegramAdapter) SendMessage(ctx context.Context, message string) error {
	msg := tgbotapi.NewMessage(t.chatID, message)

	_, err := t.bot.Send(msg)

	return err
}
