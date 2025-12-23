package handler

import (
	"fmt"
	"strings"

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
	replyText := fmt.Sprintf("Вы написали: %s", text)

	if strings.Contains(msg.Text, "подпис") {
		reply := tgbotapi.NewMessage(chatID, "О, опять про подписку? Денежки на орехи скопил? 😏\nНапиши мне по этому поводу в телеграм: @olegnastyle	")
		_, err := bot.Send(reply)
		return err
	} else {
		// Простой эхо-ответ
		reply := tgbotapi.NewMessage(chatID, replyText)
		_, err := bot.Send(reply)
		return err
	}

}
