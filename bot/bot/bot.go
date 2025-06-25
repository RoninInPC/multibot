package bot

import (
	"multibot/bot/button"
	"multibot/bot/entity"
)

type Bot interface {
	GetChannel() chan<- entity.Message
	GetType() entity.TypeBot
	GetFunctionalBuilder() button.ButtonInlineBuilder
	SetFunctionalWithStart(text string, function entity.UpdateFunc, content *button.ButtonsContent)
	Work()
}
