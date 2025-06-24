package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"multibot/bot/bot"
	"multibot/bot/button"
	"multibot/bot/button/telegrambuttons"
	"multibot/bot/update"
)

type TelegramBot struct {
	bot           *tgbotapi.BotAPI
	functional    *telegrambuttons.TelegramInlineButtons
	content       *button.ButtonsContent
	startReaction func(update update.Update)
}

func (t TelegramBot) GetType() bot.TypeBot {
	return bot.Telegram
}

func (t TelegramBot) GetFunctionalBuilder() button.ButtonInlineBuilder {
	return t.functional
}

func (t TelegramBot) SetFunctionalWithStart(text string, function func(update update.Update), content *button.ButtonsContent) {
	t.content = content

}

func (t TelegramBot) Work() {
	//TODO implement me
	panic("implement me")
}
