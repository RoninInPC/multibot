package bot

import (
	"multibot/bot/button"
	"multibot/bot/entity"
	"multibot/bot/typebot"
)

type Bot interface {
	GetChannel() chan<- entity.Message
	GetType() typebot.TypeBot
	GetFunctionalBuilder() button.ButtonInlineBuilder
	SetFunctionalWithStart(text string, function entity.UpdateFunc, content *button.ButtonsContent)
	Work()
}
