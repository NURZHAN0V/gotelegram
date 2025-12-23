package handler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CallbackHandler обрабатывает callback-запросы от инлайн-кнопок
type CallbackHandler struct{}

// NewCallbackHandler создаёт новый обработчик callback-запросов
func NewCallbackHandler() *CallbackHandler {
	return &CallbackHandler{}
}

// Handle обрабатывает callback-запрос
func (h *CallbackHandler) Handle(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) error {
	chatID := callback.Message.Chat.ID
	callbackData := callback.Data

	// Отвечаем на callback-запрос (убираем индикатор загрузки)
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := bot.Request(callbackConfig); err != nil {
		return fmt.Errorf("ошибка ответа на callback: %w", err)
	}

	// Обрабатываем различные типы callback-данных
	var replyText string
	switch callbackData {
	case "lang_ru":
		replyText = "✅ Выбран язык: Русский"
	case "lang_en":
		replyText = "✅ Выбран язык: English"
	default:
		replyText = "Неизвестная команда"
	}

	// Отправляем ответ пользователю
	reply := tgbotapi.NewMessage(chatID, replyText)
	_, err := bot.Send(reply)
	return err
}

