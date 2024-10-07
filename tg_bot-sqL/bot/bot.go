package bot

import (
	"log"
    "telegram-bot-postgres-example/database"
    "github.com/mymmrac/telego"
    "github.com/mymmrac/telego/telegoutil"
)

func StartBot() error {
	bot, err := telego.NewBot("7639067110:AAGNdt2TGgoedxh9_0I8LlV0VBYZcvYyQlk")
	if err != nil {
		return nil
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	for update := range updates {
		if update.Message != nil {
			userID := update.Message.Chat.ID
			text := update.Message.Text


			if text == "/start" {
				msg := telegoutil.Message(telegoutil.ID(userID), "Вітаю! Введіть будь-яке повідомлення, і я його збережу.")
				bot.SendMessage(msg)
			} else if text == "/my_messages" {
				messages, err := database.GetMessages(userID)
				if err != nil {
					log.Printf("Помилка отримання повідомлень: %v", err)
                    continue
				}

				response := "Ваші повідомлення:\n"
				for _, msg := range messages {
					response += "- " + msg + "\n"
				}

				msg := telegoutil.Message(telegoutil.ID(userID), response)
				bot.SendMessage(msg)
			} else {
				err := database.SaveMessage(userID, text)
				if err != nil {
					log.Printf("Помилка збереження повідомлення: %v", err)
				} else {
					msg := telegoutil.Message(telegoutil.ID(userID), "Ваше повідомлення збережено!")
					bot.SendMessage(msg)
				}
			}
		}
	}

	return nil
}