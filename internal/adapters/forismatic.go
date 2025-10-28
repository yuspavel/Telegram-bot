package adapters

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"telegram-bot/internal/entities"
	"time"
)

type ForismaticApi struct {
	client *http.Client
}

func NewForismaticApi() *ForismaticApi {
	return &ForismaticApi{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (f *ForismaticApi) GetRandomQuote(ctx context.Context) (*entities.Quote, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=ru", nil)
	if err != nil {
		return nil, errors.New("ошибка создания запроса")
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, errors.New("ошибка запроса к API")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API вернул статус не равный 200")
	}

	var qouteResp struct {
		QuoteText   string `json:"quoteText"`
		QuoteAuthor string `json:"quoteAuthor"`
		SenderName  string `json:"senderName"`
		SenderLink  string `jsnon:"senderLink"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&qouteResp); err != nil {
		return nil, errors.New("ошибка декодирования JSON")
	}

	if qouteResp.QuoteText == "" {
		return nil, errors.New("получена пустая заметка")
	}

	author := qouteResp.QuoteAuthor
	if author == "" {
		author = "Неизвестный автор"
	}

	return &entities.Quote{Text: qouteResp.QuoteText, Author: author}, nil
}
