package telegram

import (
	"YesNoTGBot/pkg/telegram/commands"
	"YesNoTGBot/pkg/telegram/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"regexp"
	"strings"
)

var commandsList = []types.Command{
	commands.HelpCommand,
	commands.AskCommand,
}

func Start(token string) {
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

		for _, cmd := range commandsList {
			re := regexp.MustCompile(strings.ToLower(cmd.Regex))
			if re.MatchString(strings.ToLower(msg.Text)) {
				args := re.FindStringSubmatch(strings.ToLower(msg.Text))

				ctx := types.Context{
					From:    msg.From,
					Chat:    msg.Chat,
					Message: msg,
					Reply: func(text string) (tgbotapi.Message, error) {
						return sendMessage(bot, msg.Chat.ID, text)
					},
				}

				cmd.Handler(&ctx, args)
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

//func sendMessageWithPhoto(bot *tgbotapi.BotAPI, chatID int64, text string, photoLink string) (tgbotapi.Message, error) {
//	msg := tgbotapi.NewPhoto(chatID, text)
//	msg.ParseMode = tgbotapi.ModeHTML
//	msg.Caption = text
//
//	bot.UploadFiles()
//
//	return bot.Send(msg)
//}
