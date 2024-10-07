package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn


func IninDB() error {
	var err error

	conn, err = pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/telegram_bot_db")
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS messages (
        id SERIAL PRIMARY KEY,
        user_id BIGINT,
        message TEXT
    );`
		_, err = conn.Exec(context.Background(), createTableSQL)
		if err != nil {
			return err
		}

		log.Println("База даних успішно ініціалізована.")
    return nil
}

func SaveMessage(userID int64, message string) error {
	query := `INSERT INTO messages (user_id, message) VALUES ($1, $2)`
	_, err := conn.Exec(context.Background(), query, userID, message)
	return err
}


func GetMessages(userID int64) ([]string, error) {
	query := `SELECT message FROM messages WHERE user_id = $1`
	rows, err := conn.Query(context.Background(), query, userID)
  if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var message string
		err = rows.Scan(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
