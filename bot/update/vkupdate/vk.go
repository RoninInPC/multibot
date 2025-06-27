package vkupdate

import (
	"github.com/SevereCloud/vksdk/v3/events"
	"multibot/bot/typebot"
)

type VKUpdate struct {
	UpdateEvent *events.MessageEventObject
	UpdateNew   *events.MessageNewObject
}

func (V VKUpdate) GetIdUserFrom() int64 {
	if V.UpdateNew != nil {
		return int64(V.UpdateNew.Message.FromID)
	}
	return int64(V.UpdateEvent.UserID)
}

func (V VKUpdate) GetType() typebot.TypeBot {
	return typebot.VK
}
