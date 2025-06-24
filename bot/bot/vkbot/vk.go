package vkbot

import (
	"multibot/bot/bot"
	"multibot/bot/button"
	"multibot/bot/update"
)

type VKBot struct {
}

func (V VKBot) GetType() bot.TypeBot {
	//TODO implement me
	panic("implement me")
}

func (V VKBot) GetFunctionalBuilder() button.ButtonInlineBuilder {
	//TODO implement me
	panic("implement me")
}

func (V VKBot) SetFunctionalWithStart(text string, function func(update update.Update), content *button.ButtonsContent) {
	//TODO implement me
	panic("implement me")
}

func (V VKBot) Work() {
	//TODO implement me
	panic("implement me")
}
