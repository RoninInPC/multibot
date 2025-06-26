package entity

import "multibot/bot/update"

type Message struct {
	Update  update.Update
	WhoID   int64
	Text    string
	Buttons interface{}
	Error   error
}

type UpdateFunc func(update update.Update, channel chan<- Message)

type TypeBot int

const (
	Telegram TypeBot = 0
	VK       TypeBot = 1
	Max      TypeBot = 2
)
