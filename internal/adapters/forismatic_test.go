package adapters

import (
	"context"
	"testing"
)

func TestNewForismaticAPI(t *testing.T) {
	api := NewForismaticApi()

	if api == nil {
		t.Error("Expecteg Forismatic API, got nil")
	}

	if api.client == nil {
		t.Errorf("Expected api.Client, got nil")
	}
}

func TestGetRandomQuote(t *testing.T) {
	api := NewForismaticApi()
	if api == nil {
		t.Error("Expecteg Forismatic API, got nil")
	}
	if api.client == nil {
		t.Errorf("Expected api.Client, got nil")
	}

	ctx := context.Background()

	quote, err := api.GetRandomQuote(ctx)
	if err != nil {
		t.Logf("API unavailable %v", err)
		return
	}

	if quote == nil {
		t.Error("Expected Quote, got nil")
		return
	}

	if quote.Text == "" {
		t.Error("Expected text, got empty text")
	}

	if quote.Author == "" {
		t.Error("Expected author, got empty author")
	}

	t.Logf("Получена цитата %s, автора %s", quote.Text, quote.Author)
}
