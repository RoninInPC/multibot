package telegramupdate

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"multibot/bot/entity"
)

type TelegramUpdate struct {
	Update *tgbotapi.Update
}

func (t TelegramUpdate) GetType() entity.TypeBot {
	return entity.Telegram
}
