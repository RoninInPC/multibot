package vkupdate

import (
	"github.com/SevereCloud/vksdk/v3/events"
)

type VKUpdate struct {
	UpdateEvent *events.MessageEventObject
	UpdateNew   *events.MessageNewObject
}
