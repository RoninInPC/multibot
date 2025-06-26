package vkbot

import (
	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
	"log"
	"multibot/bot/button"
	"multibot/bot/button/vkbuttons"
	"multibot/bot/entity"
	"multibot/bot/update/vkupdate"
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

func (V *VKBot) GetType() entity.TypeBot {
	return entity.VK
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
				if msg.Update.GetType() == entity.VK {
					update := msg.Update.(vkupdate.VKUpdate)
					if update.UpdateNew != nil {
						id = int64(update.UpdateNew.Message.FromID)
					} else {
						id = int64(update.UpdateEvent.UserID)
					}
				}
			}
			b := params.NewMessagesSendBuilder()
			b.Message("12").
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
}
