package models

import "maurodesouza/first-golang-api/database"

func Delete(id int64) (int64, error) {
	connection, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer connection.Close()

	response, err := connection.Exec(`DELETE FROM todos WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
