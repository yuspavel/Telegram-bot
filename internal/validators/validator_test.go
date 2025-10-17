package validators

import (
	"testing"
)

const (
	token      = "123456789:AbCdEfGhIjKlMoPq"
	smallToken = "123:AbCdEf"
	longToken  = "123456789:AbCdEfGhIjKlMoPqAbCdEfGhIjKlfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPqAbCdEfGhIjKlMoPq"
)

func TestValidateQoute(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		author    string
		expectErr bool
	}{
		{
			name:      "Success",
			text:      "To be or not to be",
			author:    "Sheakspeare",
			expectErr: false,
		},
		{
			name:      "Empty text",
			text:      "",
			author:    "Sheakspeare",
			expectErr: true,
		},
		{
			name:      "Empty author",
			text:      "To be or not to be",
			author:    "",
			expectErr: true,
		},
		{
			name:      "Too long text",
			text:      string(make([]byte, 1001)),
			author:    "Sheakspere",
			expectErr: true,
		},
		{
			name:      "Too long author",
			text:      "To be or not to be",
			author:    string(make([]byte, 101)),
			expectErr: true,
		},
		{
			name:      "Whitespace only text",
			text:      "   ",
			author:    "Steve Jobs",
			expectErr: true,
		},
		{
			name:      "Whitespace only author",
			text:      "To be or not to be",
			author:    "   ",
			expectErr: true,
		},
		{
			name:      "Dangerous text chars",
			text:      "To be or onclick= not to be",
			author:    "Sheakspere",
			expectErr: true,
		},
		{
			name:      "Dangerous author chars",
			text:      "To be or not to be",
			author:    "William javascript: Sheakspere",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := ValidateQuote(tt.text, tt.author)
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

func TestValidateToken(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		excpectErr bool
	}{
		{
			name:       "Success",
			token:      token,
			excpectErr: false,
		},
		{
			name:       "Empty token",
			token:      "",
			excpectErr: true,
		},
		{
			name:       "Token without :",
			token:      "1234567890AbCdEfGhIjKlMnOpQr",
			excpectErr: true,
		},
		{
			name:       "Too small token",
			token:      smallToken,
			excpectErr: true,
		},
		{
			name:       "Too long token",
			token:      longToken,
			excpectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateToken(tt.token)

			if tt.excpectErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Error("Expected nil, got error")
				}
			}
		})
	}
}

func TestValidateChatId(t *testing.T) {
	tests := []struct {
		name      string
		chatId    int64
		expectErr bool
	}{
		{
			name:      "Success",
			chatId:    123,
			expectErr: false,
		},
		{
			name:      "Chat ID too negative",
			chatId:    -1999999999999999,
			expectErr: true,
		},
		{
			name:      "Chant ID too positive",
			chatId:    1999999999999999,
			expectErr: true,
		},
		{
			name:      "Zero chat ID",
			chatId:    0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateChatID(tt.chatId)
			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Error("Expected nil, got error")
				}
			}
		})
	}
}
