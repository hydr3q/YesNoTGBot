package commands

import "YesNoTGBot/pkg/telegram/types"

var HelpCommand = types.Command{
	Regex: "^(?:/|)start(?:(.*)|)$",
	Handler: func(ctx *types.Context, args []string) {
		ctx.Reply("<b>Hello!</b>")
	},
}
