package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"multibot/bot/button"
	"multibot/bot/button/telegrambuttons"
	"multibot/bot/entity"
	"multibot/bot/typebot"
	"multibot/bot/update/telegramupdate"
)

type TelegramBot struct {
	bot            *tgbotapi.BotAPI
	functional     *telegrambuttons.TelegramInlineButtons
	handler        map[string]entity.UpdateFunc
	handlerButtons map[string]interface{}
	handlerText    map[string]string
	channel        chan entity.Message
	start          struct {
		StartCommand string
		Description  string
		Timeout      int
		Text         string
		Func         entity.UpdateFunc
		Content      interface{}
	}
	parseMod string
}

func InitTelegramBot(token, startCommandName, descriptionCommand, parseMod string, timeout int) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	t := &TelegramBot{
		bot:        bot,
		functional: telegrambuttons.InitTelegramInlineButtons(),
	}
	t.start.StartCommand = startCommandName
	t.start.Description = descriptionCommand
	t.start.Timeout = timeout
	t.parseMod = parseMod

	t.channel = make(chan entity.Message)
	t.handlerText = make(map[string]string)
	t.handler = make(map[string]entity.UpdateFunc)
	t.handlerButtons = make(map[string]interface{})

	return t, nil
}

func (t *TelegramBot) GetChannel() chan<- entity.Message {
	return t.channel
}

func (t *TelegramBot) SetFunctionalWithStart(text string, function entity.UpdateFunc, content *button.ButtonsContent) {
	t.start.Text = text
	t.start.Func = function
	t.start.Content = content.Content
	t.handler = content.Handler
	t.handlerButtons = content.HandlerButtons
	t.handlerText = content.HandlerText
}

func (t *TelegramBot) GetType() typebot.TypeBot {
	return typebot.Telegram
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
	go func() {
		for msg := range t.channel {
			var id int64
			if msg.Update == nil {
				id = msg.WhoID
			} else {
				if msg.Update.GetType() == typebot.Telegram {
					update := msg.Update.(telegramupdate.TelegramUpdate)
					if update.Update.CallbackQuery != nil {
						id = update.Update.CallbackQuery.From.ID
					} else {
						id = update.Update.Message.From.ID
					}
				}
			}
			sendMsg := tgbotapi.NewMessage(id, msg.Text)
			sendMsg.ParseMode = t.parseMod
			if msg.Buttons != nil {
				sendMsg.ReplyMarkup = msg.Buttons
			}
			_, err := t.bot.Send(sendMsg)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	for updt := range t.bot.GetUpdatesChan(u) {
		if updt.CallbackQuery != nil {
			if msg, ok := t.handler[updt.CallbackQuery.Data]; ok {
				msg(telegramupdate.TelegramUpdate{Update: &updt}, t.GetChannel())
			}
			buttons, ok := t.handlerButtons[updt.CallbackQuery.Data]
			if !ok {
				buttons = nil
			}
			text := t.handlerText[updt.CallbackQuery.Data]

			t.GetChannel() <- entity.Message{
				Update:  telegramupdate.TelegramUpdate{Update: &updt},
				Text:    text,
				Buttons: buttons}
			t.bot.Send(tgbotapi.NewDeleteMessage(updt.CallbackQuery.From.ID, updt.CallbackQuery.Message.MessageID))
		} else if updt.Message != nil {
			if updt.Message.Text == t.start.StartCommand {
				if t.start.Func != nil {
					t.start.Func(telegramupdate.TelegramUpdate{Update: &updt}, t.GetChannel())
				}
				t.GetChannel() <- entity.Message{
					Update:  telegramupdate.TelegramUpdate{Update: &updt},
					Text:    t.start.Text,
					Buttons: t.start.Content}
				t.bot.Send(tgbotapi.NewDeleteMessage(updt.Message.From.ID, updt.Message.MessageID))
			}
		}

	}
}
