package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"telegram-bot/internal/config"
	"telegram-bot/internal/handler"
	"telegram-bot/internal/middleware"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	// Создаём экземпляр бота
	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		log.Fatal("Ошибка создания бота:", err)
	}

	bot.Debug = cfg.Bot.Debug
	log.Printf("Авторизован как %s", bot.Self.UserName)

	// Создаём диспетчер обработчиков
	dispatcher := handler.NewDispatcher()

	// Регистрируем обработчики команд
	dispatcher.Register(handler.NewStartHandler())
	dispatcher.Register(handler.NewHelpHandler())
	dispatcher.Register(handler.NewInfoHandler())
	dispatcher.Register(handler.NewAdminHandler(cfg.Bot.AdminIDs))

	// Создаём обработчик обычных сообщений
	messageHandler := handler.NewMessageHandler()

	// Настраиваем получение обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = cfg.Bot.Timeout
	updates := bot.GetUpdatesChan(u)

	// Обрабатываем обновления
	for update := range updates {
		handleUpdate(bot, dispatcher, messageHandler, update)
	}
}

// handleUpdate обрабатывает одно обновление
func handleUpdate(
	bot *tgbotapi.BotAPI,
	dispatcher *handler.Dispatcher,
	messageHandler *handler.MessageHandler,
	update tgbotapi.Update,
) {
	if update.Message == nil {
		return
	}

	msg := update.Message

	if msg.IsCommand() {
		middleware.LogCommand(msg) // Логируем команду
		err := dispatcher.HandleCommand(bot, msg)
		if err != nil {
			log.Printf("Ошибка обработки команды: %v", err)
		}
		return
	}

	if msg.Text != "" {
		middleware.LogMessage(msg) // Логируем сообщение
		err := messageHandler.Handle(bot, msg)
		if err != nil {
			log.Printf("Ошибка обработки сообщения: %v", err)
		}
	}
}
