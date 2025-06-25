package telegrambuttons

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"multibot/bot/button"
	"multibot/bot/entity"
)

type TelegramInlineButtons struct {
	keyboard       [][]tgbotapi.InlineKeyboardButton
	lastKeyboard   int
	handler        map[string]entity.UpdateFunc
	handlerButtons map[string]interface{}
	handlerText    map[string]string
}

func InitTelegramInlineButtons() TelegramInlineButtons {
	return TelegramInlineButtons{
		keyboard:       make([][]tgbotapi.InlineKeyboardButton, 1),
		lastKeyboard:   0,
		handler:        make(map[string]entity.UpdateFunc),
		handlerButtons: make(map[string]interface{}),
		handlerText:    make(map[string]string),
	}
}

func (t TelegramInlineButtons) AddRow() button.ButtonInlineBuilder {
	t.lastKeyboard++
	t.keyboard = append(t.keyboard, []tgbotapi.InlineKeyboardButton{})
	return t
}

func (t TelegramInlineButtons) AddCallBack(label, text, handler string, function entity.UpdateFunc) button.ButtonInlineBuilder {
	t.keyboard[t.lastKeyboard] = append(t.keyboard[t.lastKeyboard], tgbotapi.NewInlineKeyboardButtonData(label, handler))
	t.handler[handler] = function
	t.handlerText[handler] = text
	return t
}

func (t TelegramInlineButtons) AddLink(label, url string) button.ButtonInlineBuilder {
	t.keyboard[t.lastKeyboard] = append(t.keyboard[t.lastKeyboard], tgbotapi.NewInlineKeyboardButtonURL(label, url))
	return t
}

func (t TelegramInlineButtons) AddText(label, text, handler string) button.ButtonInlineBuilder {
	t.keyboard[t.lastKeyboard] = append(t.keyboard[t.lastKeyboard], tgbotapi.NewInlineKeyboardButtonData(label, handler))
	t.handlerText[handler] = text
	return t
}

func (t TelegramInlineButtons) AddButtonsCallback(label, text, handler string, buttons *button.ButtonsContent) button.ButtonInlineBuilder {
	t.keyboard[t.lastKeyboard] = append(t.keyboard[t.lastKeyboard], tgbotapi.NewInlineKeyboardButtonData(label, handler))
	t.handlerText[handler] = text
	t.handlerButtons[handler] = buttons.Content
	for k, v := range t.handlerButtons {
		t.handlerButtons[k] = v
	}
	for k, v := range t.handler {
		t.handler[k] = v
	}
	for k, v := range t.handlerText {
		t.handlerText[k] = v
	}
	return t
}

func (t TelegramInlineButtons) Build() *button.ButtonsContent {
	return &button.ButtonsContent{
		Content:        tgbotapi.NewInlineKeyboardMarkup(t.keyboard...),
		Handler:        t.handler,
		HandlerButtons: t.handlerButtons,
		HandlerText:    t.handlerText,
	}
}
