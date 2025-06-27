package button

import (
	"multibot/bot/entity"
)

type ButtonsContent struct {
	Content        interface{}
	Handler        map[string]entity.UpdateFunc
	HandlerButtons map[string]interface{}
	HandlerText    map[string]string
}

type ButtonInlineBuilder interface {
	AddRow() ButtonInlineBuilder
	AddCallBack(label, handler string, function entity.UpdateFunc) ButtonInlineBuilder
	AddLink(label, url string) ButtonInlineBuilder
	AddText(label, text, handler string) ButtonInlineBuilder
	AddButtonsCallback(label, text, handler string, buttons *ButtonsContent) ButtonInlineBuilder
	Build() *ButtonsContent
	GetNewBuilder() ButtonInlineBuilder
}
