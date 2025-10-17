package config

import (
	"log/slog"
	"os"
	"testing"
)

const (
	token  = "123456789:AbCdEfGhIjKlMoPq"
	chatId = "123456789"
)

func TestLoadConfig(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	tests := []struct {
		name              string
		botToken          string
		chatID            string
		sendTestQuote     string
		errExpected       bool
		expectedTestQuote bool
	}{
		{
			name:              "Valid config with default test quote",
			botToken:          token,
			chatID:            chatId,
			sendTestQuote:     "",
			errExpected:       false,
			expectedTestQuote: true,
		},
		{
			name:              "Valid config with eenabled test quote",
			botToken:          token,
			chatID:            chatId,
			sendTestQuote:     "true",
			errExpected:       false,
			expectedTestQuote: true,
		},
		{
			name:              "Invalid config with unknown test quote",
			botToken:          token,
			chatID:            chatId,
			sendTestQuote:     "false",
			errExpected:       false,
			expectedTestQuote: false,
		},
		{
			name:          "Invalid config with empty botToken",
			botToken:      "",
			chatID:        chatId,
			sendTestQuote: "",
			errExpected:   true,
		},
		{
			name:          "Invalid config with empty chatID",
			botToken:      token,
			chatID:        "",
			sendTestQuote: "",
			errExpected:   true,
		},
		{
			name:     "Invalid chatID",
			botToken: token,
			chatID:   "abc",
			//sendTestQuote: "",
			errExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originBoToken := tt.botToken
			originChatID := tt.chatID
			originSendTestQuote := tt.sendTestQuote

			os.Setenv("BOT_TOKEN", tt.botToken)
			os.Setenv("CHAT_ID", tt.chatID)
			os.Setenv("SEND_TEST_QUOTE", tt.sendTestQuote)

			defer func() {
				os.Setenv("BOT_TOKEN", originBoToken)
				os.Setenv("CHAT_ID", originChatID)
				os.Setenv("SEND_TEST_QUOTE", originSendTestQuote)
			}()

			config, err := LoadConfig(logger)

			if tt.errExpected {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected nil, got error %s", err.Error())
				}
				if config == nil {
					t.Error("Expected config, got nil")
				} else {
					if tt.botToken != config.BotToken {
						t.Errorf("Expected %s token, got %s", tt.botToken, config.BotToken)
					}
					if config.ChatId != 123456789 {
						t.Errorf("Expected %s chatID, got %d", chatId, config.ChatId)
					}
					if config.SendTestQuote != tt.expectedTestQuote {
						t.Errorf("Expected SendTestQoute=%t, got %t", config.SendTestQuote, tt.expectedTestQuote)
					}
				}
			}

		})
	}
}

func TestMaskToken(t *testing.T) {
	/*logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))*/

	tests := []struct {
		name     string
		token    string
		expected string
	}{
		{
			name:     "Empty token",
			token:    "",
			expected: "***",
		},
		{
			name:     "Token size less than 8",
			token:    "123:abc",
			expected: "***",
		},
		{
			name:     "OK",
			token:    "123456789:abcdefghijklmop",
			expected: "1234***lmop",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			str := maskToken(tt.token)

			if str != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, str)
			}

		})
	}
}
