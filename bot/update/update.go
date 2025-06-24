package update

import "multibot/bot/bot"

type Info map[string]interface{}

type Update interface {
	GetType() bot.TypeBot
}
