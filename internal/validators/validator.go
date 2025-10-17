package validators

import (
	"errors"
	"strings"
)

func ValidateQuote(text, author string) error {
	if strings.TrimSpace(text) == "" {
		return errors.New("текст цитаты не может быть пустым")
	}

	if strings.TrimSpace(author) == "" {
		return errors.New("автор цитаты не может быть пустым")
	}

	if len(text) > 1000 {
		return errors.New("заметка слишком длинная. Максимум 1000 символов")
	}

	if len(author) > 100 {
		return errors.New("имя автора слишком длинное. Максимум 100 символов")
	}

	if containsDangerousChars(text) {
		return errors.New("текст цитаты не должен содержать запретные слова")
	}

	if containsDangerousChars(author) {
		return errors.New("имя автора не должно содержать запрещенные слова")
	}
	return nil
}

func ValidateToken(token string) error {
	if strings.TrimSpace(token) == "" {
		return errors.New("токен не может быть пустым")
	}

	if !strings.Contains(token, ":") {
		return errors.New("неверный формат токена")
	}

	if len(token) < 20 || len(token) > 100 {
		return errors.New("неверная длина токена")
	}
	return nil
}

func ValidateChatID(chatId int64) error {
	if chatId == 0 {
		return errors.New("id чата не может быть равным нулю")
	}

	if chatId < -999999999999999 || chatId > 999999999999999 {
		return errors.New("id чата вне допустимого диапазона")
	}
	return nil
}

func containsDangerousChars(text string) bool {
	dangerousChars := []string{"<script", "</script>", "javascript:", "data:",
		"vbscript:", "onload=", "onerror=", "onclick="}

	textLower := strings.ToLower(text)
	for _, char := range dangerousChars {
		if strings.Contains(textLower, char) {
			return true
		}
	}
	return false
}
