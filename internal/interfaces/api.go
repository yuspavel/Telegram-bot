package interfaces

import (
	"context"
	"telegram-bot/internal/entities"
)

type QuoteAPI interface {
	GetRandomQuote(ctx context.Context) (*entities.Quote, error)
}

type TelegramSender interface {
	SendMessage(ctx context.Context, message string) error
}

type CronScheduler interface {
	Start()
	AddJob(spec string, job func())
}
