package models

import (
	"maurodesouza/first-golang-api/database"
)

func Update(id int64, todo Todo) (int64, error) {
	connection, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer connection.Close()

	response, err := connection.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
