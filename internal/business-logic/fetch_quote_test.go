package businesslogic

import (
	"context"
	"errors"
	"telegram-bot/internal/entities"
	"testing"
)

type QuoteMockApi struct {
	quote *entities.Quote
	err   error
}

func (qa *QuoteMockApi) GetRandomQuote(ctx context.Context) (*entities.Quote, error) {
	return qa.quote, qa.err
}

func TestFetchQuoteService(t *testing.T) {
	tests := []struct {
		name          string
		mock          *QuoteMockApi
		expectedErr   bool
		expectedQuote *entities.Quote
	}{

		{
			name: "Success",
			mock: &QuoteMockApi{
				quote: &entities.Quote{
					Text:   "Test text",
					Author: "Test author",
				},
				err: nil,
			},
			expectedErr:   false,
			expectedQuote: &entities.Quote{Text: "Test text", Author: "Test author"},
		},
		{
			name: "Nil quote",
			mock: &QuoteMockApi{
				quote: nil,
				err:   nil,
			},
			expectedErr:   false,
			expectedQuote: nil,
		},
		{
			name: "API error",
			mock: &QuoteMockApi{
				quote: nil,
				err:   errors.New("API error"),
			},
			expectedErr:   true,
			expectedQuote: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fqc := NewFetchQuoteService(tt.mock)
			quote, err := fqc.FetchQuote(context.Background())

			if tt.expectedErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected nil, got error %s", err.Error())
				}
				if tt.expectedQuote != nil {
					if quote == nil {
						t.Error("Expected quote, got nil")
					} else {
						if tt.expectedQuote.Text != quote.Text {
							t.Errorf("Expected text %s, got %s", tt.expectedQuote.Text, quote.Text)
						}
						if tt.expectedQuote.Author != quote.Author {
							t.Errorf("Expected author %s got %s", tt.expectedQuote.Author, quote.Author)
						}
					}
				}
			}
		})
	}
}
