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

func Delete(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", id)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	rows, err := models.Delete(int64(id))

	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %v registros", rows)
	}

	var resp map[string]any

	resp = map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo removido com sucesso: %v", id),
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(resp)
}
