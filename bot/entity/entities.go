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
