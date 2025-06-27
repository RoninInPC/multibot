package main

import (
	"fmt"
	"multibot/bot/bot/vkbot"
	"multibot/bot/entity"
	"multibot/bot/update"
)

func main() {
	bot, err := /*telegrambot.InitTelegramBot(
		"7676691783:AAGLUFpXSt_QwWgArH44lUT5xaQrudB4aTc",
		"/start",
		"начнём?",
		tgbotapi.ModeMarkdown,
		40)*/
		vkbot.InitVKBot(
			"vk1.a.nPK8pGqNdwsaWVOvnXDvQuoSBy3zAjrlPt6J3GDlixg7S4j8q_nAIEQFPW9Ns5Bi2-QuargkpKbln8Rf5sBZbByvxF0ac56IhOf9P0bt7JmHKLB35elEZT1T3cnIHpW6nhWpesNHgbU9SkPGk9o5_ZHSreIWRzTpb67tAXMobXDF6aNPp2dtT3kCsoHxrPRiEuvXHqaEMoyM3wmUO4qLWA",
			"Начать",
			231066346,
		)
	if err != nil {
		panic(err)
	}
	builder := bot.GetFunctionalBuilder()
	builder.AddText("Кнопочка", "Привет от кнопочки", "1").
		AddCallBack("Узнать id", "2",
			func(update update.Update, channel chan<- entity.Message) {
				id := update.GetIdUserFrom()
				channel <- entity.Message{
					Update: update,
					WhoID:  id,
					Text:   fmt.Sprintf("Ваш id равен %d и ,благодаря ему, к вам пришло сообщение", id),
				}
			}).
		AddCallBack("Отправить Сергею", "3", func(update update.Update, channel chan<- entity.Message) {
			var id int64 = 287184
			channel <- entity.Message{
				Update: nil,
				WhoID:  id,
				Text:   fmt.Sprintf("Тестирование отправки вне сообщества по id,id равен %d и ,благодаря ему, к вам пришло сообщение", id),
			}
		})
	bot.SetFunctionalWithStart("Добрый день!", nil, builder.Build())
	bot.Work()
}
