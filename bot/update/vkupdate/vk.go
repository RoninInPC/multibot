package vkupdate

import (
	"github.com/SevereCloud/vksdk/v3/events"
	"multibot/bot/entity"
)

type VKUpdate struct {
	UpdateEvent *events.MessageEventObject
	UpdateNew   *events.MessageNewObject
}

func (V VKUpdate) GetType() entity.TypeBot {
	return entity.VK
}
