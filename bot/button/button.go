package button

import (
	"multibot/bot/update"
)

type ButtonsContent struct {
	Content        interface{}
	Handler        map[string]func(update update.Update)
	HandlerButtons map[string]interface{}
	HandlerText    map[string]string
	HandlerDelete  map[string]bool
}

type ButtonInlineBuilder interface {
	AddRow() ButtonInlineBuilder
	AddCallBack(label, text, handler string, function func(update update.Update)) ButtonInlineBuilder
	AddLink(label, url string) ButtonInlineBuilder
	AddText(label, text, handler string) ButtonInlineBuilder
	AddButtonsCallback(label, text, handler string, buttons *ButtonsContent) ButtonInlineBuilder
	Build() *ButtonsContent
}
