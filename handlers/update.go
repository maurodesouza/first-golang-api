package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"maurodesouza/first-golang-api/models"

	"github.com/go-chi/chi/v5"

	"net/http"
	"strconv"
)

func Update(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", id)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	var todo models.Todo

	err = json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	rows, err := models.Update(int64(id), todo)

	if err != nil {
		log.Printf("Erro ao atualizar todo: %v", id)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %v registros", rows)
	}

	var resp map[string]any

	resp = map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo atualizado com sucesso: %v", id),
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(resp)
}
