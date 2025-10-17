package businesslogic

import (
	"context"
	"errors"
	"telegram-bot/internal/entities"
	"telegram-bot/internal/interfaces"
	"telegram-bot/internal/validators"
)

type FetchQuoteService struct {
	api interfaces.QuoteAPI
}

func NewFetchQuoteService(api interfaces.QuoteAPI) *FetchQuoteService {
	return &FetchQuoteService{api: api}
}

func (f *FetchQuoteService) FetchQuote(ctx context.Context) (*entities.Quote, error) {
	quote, err := f.api.GetRandomQuote(ctx)
	if err != nil {
		return nil, errors.New("не удалось получить цитату")
	}

	if quote != nil {
		if err := validators.ValidateQuote(quote.Text, quote.Author); err != nil {
			return nil, errors.New("получена невалидная цитата" + err.Error())
		}
	}
	return quote, nil
}
