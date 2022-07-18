package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"maurodesouza/first-golang-api/models"
	"net/http"
)

func Create(response http.ResponseWriter, request *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso: %v", id),
		}
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(resp)
}
