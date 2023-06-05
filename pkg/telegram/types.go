package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type command struct {
	regex   string
	handler func(*Context, []string)
}

type Context struct {
	from    *tgbotapi.User
	chat    *tgbotapi.Chat
	message *tgbotapi.Message
	reply   func(text string) (tgbotapi.Message, error)
}
