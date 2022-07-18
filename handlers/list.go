package handlers

import (
	"encoding/json"
	"log"
	"maurodesouza/first-golang-api/models"
	"net/http"
)

func List(response http.ResponseWriter, request *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Erro ao obter todos: %v", err)
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(todos)
}
