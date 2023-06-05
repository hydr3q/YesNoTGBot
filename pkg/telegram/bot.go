package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"regexp"
)

var commands = []command{}

func start(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // игнорируем пустые сообщения
			continue
		}

		msg := update.Message

		for _, cmd := range commands {
			re := regexp.MustCompile(cmd.regex)
			if re.MatchString(msg.Text) {
				args := re.FindStringSubmatch(msg.Text)

				ctx := Context{
					from:    msg.From,
					chat:    msg.Chat,
					message: msg,
					reply: func(text string) (tgbotapi.Message, error) {
						return sendMessage(bot, msg.Chat.ID, text)
					},
				}

				cmd.handler(&ctx, args)
				break
			}
		}
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML
	return bot.Send(msg)
}
