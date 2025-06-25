package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"multibot/bot/button"
	"multibot/bot/button/telegrambuttons"
	"multibot/bot/entity"
	"multibot/bot/update/telegramupdate"
)

type TelegramBot struct {
	bot            *tgbotapi.BotAPI
	functional     *telegrambuttons.TelegramInlineButtons
	Handler        map[string]entity.UpdateFunc
	HandlerButtons map[string]interface{}
	HandlerText    map[string]string
	channel        chan entity.Message
	start          struct {
		StartCommand string
		Description  string
		Timeout      int
		Text         string
		Func         entity.UpdateFunc
		Content      interface{}
	}
}

func (t *TelegramBot) GetChannel() chan<- entity.Message {
	return t.channel
}

func (t *TelegramBot) SetFunctionalWithStart(text string, function entity.UpdateFunc, content *button.ButtonsContent) {
	t.start.Text = text
	t.start.Func = function
	t.start.Content = content
	t.Handler = content.Handler
	t.HandlerButtons = content.HandlerButtons
	t.HandlerText = content.HandlerText
}

func (t *TelegramBot) GetType() entity.TypeBot {
	return entity.Telegram
}

func (t *TelegramBot) GetFunctionalBuilder() button.ButtonInlineBuilder {
	return t.functional
}

func (t *TelegramBot) Work() {
	cmdCfg := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			t.start.StartCommand,
			t.start.Description},
	)
	_, _ = t.bot.Send(cmdCfg)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = t.start.Timeout
	for updt := range t.bot.GetUpdatesChan(u) {
		if updt.CallbackQuery != nil {
			if msg, ok := t.Handler[updt.CallbackQuery.Data]; ok {
				msg(telegramupdate.TelegramUpdate{Update: &updt}, t.GetChannel())
			}
			buttons, ok := t.HandlerButtons[updt.CallbackQuery.Data]
			if !ok {
				buttons = nil
			}
			text := t.HandlerText[updt.CallbackQuery.Data]

			t.GetChannel() <- entity.Message{Text: text, Buttons: buttons}

		} else if updt.Message != nil {

		}

	}
}
