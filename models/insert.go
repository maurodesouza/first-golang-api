package models

import "maurodesouza/first-golang-api/database"

func Insert(todo Todo) (id int64, err error) {
	connection, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer connection.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = connection.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
