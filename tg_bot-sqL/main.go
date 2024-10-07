package main

import (
	"log"
  "telegram-bot-postgres-example/bot"
  "telegram-bot-postgres-example/database"
)

func main() {
	err := database.IninDB()
  if err != nil {
		log.Fatalf("Помилка ініціалізації бази даних: %v", err)
	}

	err = bot.StartBot()
	if err != nil {
		log.Fatalf("Помилка запуску бота: %v", err)
	}
}