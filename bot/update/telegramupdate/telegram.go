package telegramupdate

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"multibot/bot/typebot"
)

type TelegramUpdate struct {
	Update *tgbotapi.Update
}

func (t TelegramUpdate) GetIdUserFrom() int64 {
	if t.Update.CallbackQuery != nil {
		return t.Update.CallbackQuery.From.ID
	} else {
		return t.Update.Message.From.ID
	}
}

func (t TelegramUpdate) GetType() typebot.TypeBot {
	return typebot.Telegram
}
