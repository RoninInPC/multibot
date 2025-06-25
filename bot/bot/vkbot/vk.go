package vkbot

import (
	"multibot/bot/button"
	"multibot/bot/entity"
)

type VKBot struct {
}

func (V VKBot) GetChannel() chan<- entity.Message {
	//TODO implement me
	panic("implement me")
}

func (V VKBot) SetFunctionalWithStart(text string, function entity.UpdateFunc, content *button.ButtonsContent) {
	//TODO implement me
	panic("implement me")
}

func (V VKBot) GetType() entity.TypeBot {
	//TODO implement me
	panic("implement me")
}

func (V VKBot) GetFunctionalBuilder() button.ButtonInlineBuilder {
	//TODO implement me
	panic("implement me")
}

func (V VKBot) Work() {
	//TODO implement me
	panic("implement me")
}
