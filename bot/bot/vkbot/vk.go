package vkbot

import (
	"context"
	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/SevereCloud/vksdk/v3/events"
	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
	"log"
	"multibot/bot/button"
	"multibot/bot/button/vkbuttons"
	"multibot/bot/entity"
	"multibot/bot/typebot"
	"multibot/bot/update/vkupdate"
	"strings"
)

type VKBot struct {
	bot            *api.VK
	lp             *longpoll.LongPoll
	functional     *vkbuttons.VKInlineButtons
	channel        chan entity.Message
	groupID        int
	handler        map[string]entity.UpdateFunc
	handlerButtons map[string]interface{}
	handlerText    map[string]string
	start          struct {
		StartHandler string
		Text         string
		Func         entity.UpdateFunc
		Content      interface{}
	}
}

func InitVKBot(token, startHandler string, groupID int) (*VKBot, error) {
	bot := api.NewVK(token)
	lp, err := longpoll.NewLongPoll(bot, groupID)
	if err != nil {
		return nil, err
	}
	vk := &VKBot{
		bot:            bot,
		lp:             lp,
		functional:     vkbuttons.InitVkInlineButtons(),
		channel:        make(chan entity.Message),
		handler:        make(map[string]entity.UpdateFunc),
		handlerText:    make(map[string]string),
		handlerButtons: make(map[string]interface{}),
	}
	vk.start.StartHandler = startHandler
	return vk, nil

}

func (V *VKBot) GetChannel() chan<- entity.Message {
	return V.channel
}

func (V *VKBot) SetFunctionalWithStart(text string, function entity.UpdateFunc, content *button.ButtonsContent) {
	V.start.Text = text
	V.start.Func = function
	V.start.Content = content.Content
	V.handlerText = content.HandlerText
	V.handlerButtons = content.HandlerButtons
	V.handler = content.Handler
}

func (V *VKBot) GetType() typebot.TypeBot {
	return typebot.VK
}

func (V *VKBot) GetFunctionalBuilder() button.ButtonInlineBuilder {
	return V.functional
}

func (V *VKBot) Work() {
	go func() {
		for msg := range V.channel {
			var id int64
			if msg.Update == nil {
				id = msg.WhoID
			} else {
				if msg.Update.GetType() == typebot.VK {
					update := msg.Update.(vkupdate.VKUpdate)
					if update.UpdateNew != nil {
						id = int64(update.UpdateNew.Message.FromID)
					} else {
						id = int64(update.UpdateEvent.UserID)
					}
				}
			}
			b := params.NewMessagesSendBuilder()
			b.Message(msg.Text).
				RandomID(0).
				PeerID(int(id))
			if msg.Buttons != nil {
				b.Keyboard(msg.Buttons)
			}
			_, err := V.bot.MessagesSend(b.Params)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	V.lp.MessageEvent(func(ctx context.Context, object events.MessageEventObject) {
		payload := string(object.Payload)
		payload = strings.Replace(payload, "\"", "", -1)
		if msg, ok := V.handler[payload]; ok {
			msg(vkupdate.VKUpdate{UpdateEvent: &object}, V.GetChannel())
		}
		buttons, ok := V.handlerButtons[payload]
		if !ok {
			buttons = nil
		}
		text := V.handlerText[payload]

		V.GetChannel() <- entity.Message{
			Update:  vkupdate.VKUpdate{UpdateEvent: &object},
			Text:    text,
			Buttons: buttons,
		}
		ans := params.NewMessagesSendMessageEventAnswerBuilder().
			EventID(object.EventID).UserID(object.UserID).PeerID(object.PeerID)
		_, err := V.bot.MessagesSendMessageEventAnswer(ans.Params)
		if err != nil {
			panic(err)
		}

	})
	V.lp.MessageNew(func(ctx context.Context, object events.MessageNewObject) {
		if object.Message.Text == V.start.StartHandler {
			if V.start.Func != nil {
				V.start.Func(vkupdate.VKUpdate{UpdateNew: &object}, V.GetChannel())
			}
			V.GetChannel() <- entity.Message{
				Update:  vkupdate.VKUpdate{UpdateNew: &object},
				Text:    V.start.Text,
				Buttons: V.start.Content}
		}
	})
	V.lp.Run()
}
