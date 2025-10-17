package businesslogic

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"telegram-bot/internal/entities"
	"telegram-bot/internal/interfaces"
)

type SendQuoteService struct {
	api interfaces.TelegramSender
}

func NewSendQuoteService(api interfaces.TelegramSender) *SendQuoteService {
	return &SendQuoteService{api: api}
}

func (s *SendQuoteService) SendQuote(ctx context.Context, quote *entities.Quote) error {
	if quote == nil {
		return errors.New("цитата не может быть nil при отправке")
	}

	msg := s.FormatQuote(quote)
	err := s.api.SendMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("отправка цитаты не удалась. %s", err.Error())
	}
	return nil
}

func (s *SendQuoteService) FormatQuote(quote *entities.Quote) string {
	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	text := quote.Text
	if len(text) > 200 {
		text = text[:197] + "..."
	}

	styles := []func(string, string) string{s.formatStyle2, s.formatStyle3}
	style := styles[rand.Intn(len(styles))]
	return style(text, quote.Author)

}

func (s *SendQuoteService) formatStyle2(text, author string) string {
	emojis := []string{"💫", "✨", "🌟", "🎯", "🔥", "💡", "🌈", "🦋", "🌸", "🎪"}
	emoji := emojis[rand.Intn(len(emojis))]

	return fmt.Sprintf("%s *Мудрая мысль*\n\n"+
		"❝ %s ❞\n\n"+
		"    — *%s* ✍️", emoji, text, author)
}

func (s *SendQuoteService) formatStyle3(text, author string) string {
	emojis := []string{"💫", "✨", "🌟", "🎯", "🔥", "💡", "🌈", "🦋", "🌸", "🎪"}
	emoji := emojis[rand.Intn(len(emojis))]

	return fmt.Sprintf("\"%s *Вдохновение дня*\n\n"+
		"━━━━━━━━━━━━━━━━━━━━━━━━━━\n"+
		"  %s\n"+
		"━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n"+
		"👤 *%s*", emoji, text, author)
}
