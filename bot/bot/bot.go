package bot

import (
	"multibot/bot/button"
	"multibot/bot/update"
)

type TypeBot int

const (
	Telegram TypeBot = 0
	VK       TypeBot = 1
	Max      TypeBot = 2
)

type Bot interface {
	GetType() TypeBot
	GetFunctionalBuilder() button.ButtonInlineBuilder
	SetFunctionalWithStart(text string, function func(update update.Update), content *button.ButtonsContent)
	Work()
}
