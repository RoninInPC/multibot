package update

import (
	"multibot/bot/entity"
)

type Info map[string]interface{}

type Update interface {
	GetType() entity.TypeBot
}
