package vkbuttons

import (
	"github.com/SevereCloud/vksdk/v3/object"
	"multibot/bot/button"
	"multibot/bot/entity"
)

type VKInlineButtons struct {
	keyboard       *object.MessagesKeyboard
	handler        map[string]entity.UpdateFunc
	handlerButtons map[string]interface{}
	handlerText    map[string]string
}

func InitVkInlineButtons() *VKInlineButtons {
	return &VKInlineButtons{
		keyboard:       object.NewMessagesKeyboardInline(),
		handler:        make(map[string]entity.UpdateFunc),
		handlerButtons: make(map[string]interface{}),
		handlerText:    make(map[string]string),
	}
}

func (V VKInlineButtons) AddRow() button.ButtonInlineBuilder {
	V.keyboard.AddRow()
	return V
}

func (V VKInlineButtons) AddCallBack(label, text, handler string, function entity.UpdateFunc) button.ButtonInlineBuilder {
	V.keyboard = V.keyboard.AddCallbackButton(label, handler, "secondary")
	V.handlerText[handler] = text
	V.handler[handler] = function
	return V
}

func (V VKInlineButtons) AddLink(label, url string) button.ButtonInlineBuilder {
	V.keyboard = V.keyboard.AddOpenLinkButton(url, label, "")
	return V
}

func (V VKInlineButtons) AddText(label, text, handler string) button.ButtonInlineBuilder {
	V.keyboard = V.keyboard.AddTextButton(label, text, "secondary")
	return V
}

func (V VKInlineButtons) AddButtonsCallback(label, text, handler string, buttons *button.ButtonsContent) button.ButtonInlineBuilder {
	V.keyboard = V.keyboard.AddCallbackButton(label, handler, "secondary")
	V.handlerText[handler] = text
	V.handlerButtons[handler] = buttons.Content
	for k, v := range V.handlerButtons {
		V.handlerButtons[k] = v
	}
	for k, v := range V.handler {
		V.handler[k] = v
	}
	for k, v := range V.handlerText {
		V.handlerText[k] = v
	}
	return V
}

func (V VKInlineButtons) Build() *button.ButtonsContent {
	return &button.ButtonsContent{
		Content:        V.keyboard,
		Handler:        V.handler,
		HandlerButtons: V.handlerButtons,
		HandlerText:    V.handlerText,
	}
}
