package handler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// MessageHandler обрабатывает обычные текстовые сообщения
type MessageHandler struct{}

// NewMessageHandler создаёт новый обработчик сообщений
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

// Handle обрабатывает текстовое сообщение
func (h *MessageHandler) Handle(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	chatID := msg.Chat.ID
	text := msg.Text

	// Простой эхо-ответ
	replyText := fmt.Sprintf("Вы написали: %s", text)

	reply := tgbotapi.NewMessage(chatID, replyText)
	_, err := bot.Send(reply)
	return err
}
