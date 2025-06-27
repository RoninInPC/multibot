package update

import (
	"multibot/bot/typebot"
)

type Info map[string]interface{}

type Update interface {
	GetType() typebot.TypeBot
	GetIdUserFrom() int64
}
