package types

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Context struct {
	From    *tgbotapi.User
	Chat    *tgbotapi.Chat
	Message *tgbotapi.Message
	Reply   func(text string) (tgbotapi.Message, error)
}
