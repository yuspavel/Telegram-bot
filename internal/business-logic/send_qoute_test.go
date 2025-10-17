package businesslogic

import (
	"context"
	"errors"
	"telegram-bot/internal/entities"
	"testing"
)

type SendMockApi struct {
	err error
}

func (sma *SendMockApi) SendMessage(ctx context.Context, message string) error {
	return sma.err
}

func TestSendMessage(t *testing.T) {

	tests := []struct {
		name      string
		mock      *SendMockApi
		qoute     *entities.Quote
		expectErr bool
	}{
		{
			name:      "Success",
			mock:      &SendMockApi{err: nil},
			qoute:     &entities.Quote{Text: "Test text", Author: "Test author"},
			expectErr: false,
		},
		{
			name: "Send error",
			mock: &SendMockApi{
				err: errors.New("error"),
			},
			qoute: &entities.Quote{
				Text:   "Test text",
				Author: "Test author",
			},
			expectErr: true,
		},
		{
			name:      "Nil message",
			mock:      &SendMockApi{err: errors.New("nil message dosn't allow")},
			qoute:     nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqs := NewSendQuoteService(tt.mock)

			err := sqs.SendQuote(context.Background(), tt.qoute)

			if tt.expectErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Error("Expected nil, got error")
				}
			}
		})
	}
}
