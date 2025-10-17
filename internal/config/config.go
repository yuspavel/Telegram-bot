package config

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
	"telegram-bot/internal/validators"
)

type Config struct {
	BotToken      string
	ChatId        int64
	SendTestQuote bool
}

func LoadConfig(logger *slog.Logger) (*Config, error) {
	botToken := os.Getenv("BOT_TOKEN")
	chatIdstr := os.Getenv("CHAT_ID")
	sendTestQuoteStr := os.Getenv("SEND_TEST_QUOTE")

	if botToken == "" || chatIdstr == "" {
		logger.Error("Не найдены переменные окружения", "BOT_TOKEN", maskToken(botToken), "CHAT_ID", chatIdstr)
		return nil, errors.New("значения переменных окружения отсутствуют")
	}

	chatId, err := strconv.ParseInt(chatIdstr, 10, 64)
	if err != nil {
		logger.Error("ошибка преобразования ChatId в int64", "error", err)
		return nil, err
	}

	if err = validators.ValidateToken(botToken); err != nil {
		logger.Error("невалидный токен бота", "error", err)
		return nil, err
	}

	if err = validators.ValidateChatID(chatId); err != nil {
		logger.Error("невалидный ID чата", "error", err)
		return nil, err
	}

	sendTestQuote := true
	if sendTestQuoteStr != "" {
		if parced, err := strconv.ParseBool(sendTestQuoteStr); err == nil {
			sendTestQuote = parced
		} else {
			logger.Warn("Неверное значение переменной окружения SEND_TEST_QUOTE. Значение по умолчанию true", "value", sendTestQuoteStr)
		}
	}
	return &Config{BotToken: botToken,
			ChatId:        chatId,
			SendTestQuote: sendTestQuote},
		nil
}

func maskToken(token string) string {
	if token == "" {
		return "***"
	}
	if len(token) < 8 {
		return "***"
	}
	return token[:4] + "***" + token[len(token)-4:] //----------Первые 4 символа токена + ***+значение токена начиная с позиции len(token)-4 (последние 4 символа токена)
}
