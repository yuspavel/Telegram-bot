package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"telegram-bot/internal/adapters"
	businesslogic "telegram-bot/internal/business-logic"
	"telegram-bot/internal/config"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func SetupLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func main() {

	logger := SetupLogger()

	if err := godotenv.Load("vars.env"); err != nil { //----Функция Load (если без параметров) находит .env файл в текущей директории и создает в системе описанные в нем переменные окружения со значениями
		logger.Warn("Файл .env не найден", "error", err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM) //-----------Контекст для безопасного завершения приложения
	defer cancel()

	config, err := config.LoadConfig(logger) //---------------------Наш метод LoadConfig считывает переменные окружения из системы (те что были созданы функцией Load выше)
	if err != nil {
		logger.Error("Ошибка загрузки конфигурации", "error", err)
		return
	}

	quoteAPI := adapters.NewForismaticApi() //-----------------------------------------Создаем экземпляр адаптера (API)
	telegramAPI, err := adapters.NewTelegramAdapter(config.BotToken, config.ChatId)
	if err != nil {
		logger.Error("Невозможно создать telegram адаптер", "error", err)
		os.Exit(1)
	}

	//Соединение реализации методов и интерфейсов
	fetchQuoteService := businesslogic.NewFetchQuoteService(quoteAPI)
	sendQuteService := businesslogic.NewSendQuoteService(telegramAPI)

	c := cron.New()
	defer c.Stop()

	_, err = c.AddFunc("0 8,9,10,11,12,13,14,15,16,20,21,22 * * *", func() {

		taskCtx := context.Background()

		quote, err := fetchQuoteService.FetchQuote(taskCtx)
		if err != nil {
			logger.Error("Не удалось получить заметку", "error", err)
			return
		}

		err = sendQuteService.SendQuote(taskCtx, quote)
		if err != nil {
			logger.Error("Не удалось отправить заметку", "error", err)
		} else {
			logger.Info("Заметка отправлена успешно", "Qoute", quote.Text, "Author", quote.Author)
		}
	})
	if err != nil {
		logger.Error("Не удалось создать расписание", "error", err)
		os.Exit(1)
	}

	c.Start()
	logger.Info("Отправка тестовой цитаты")

	if config.SendTestQuote {
		ctx := context.Background()

		quote, err := fetchQuoteService.FetchQuote(ctx)
		if err != nil {
			logger.Error("Не удалось получить тестовую заметку", "error", err)
		} else {
			if err = sendQuteService.SendQuote(ctx, quote); err != nil {
				logger.Error("Не удалось отправить тестовую заметку", "error", err)
			} else {
				logger.Info("Тестовая заметка успешно отправлена")
			}
		}
	} else {
		logger.Info("Отправка тестовой заметки отключена")
	}

	<-ctx.Done()
	logger.Info("Останавливаем планировщик...")
	stopCtx := c.Stop()
	<-stopCtx.Done()
	logger.Info("Планировшик остановлен.")

}
