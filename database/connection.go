package database

import (
	"database/sql"
	"fmt"
	"maurodesouza/first-golang-api/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Database,
	)

	connection, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
