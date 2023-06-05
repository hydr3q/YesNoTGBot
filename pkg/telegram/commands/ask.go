package commands

import (
	"YesNoTGBot/pkg/api"
	"YesNoTGBot/pkg/telegram/types"
	"fmt"
)

var AskCommand = types.Command{
	Regex: "",
	Handler: func(ctx *types.Context, args []string) {
		answer := api.GetAnswer()
		_, err := ctx.Reply("Ответ: <b>" + answer.Value + "</b>\n\n" + answer.Image)
		if err != nil {
			fmt.Println(err)
		}
	},
}
