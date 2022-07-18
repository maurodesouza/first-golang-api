package handlers

import (
	"encoding/json"
	"log"
	"maurodesouza/first-golang-api/models"

	"github.com/go-chi/chi/v5"

	"net/http"
	"strconv"
)

func Get(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", id)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	todo, err := models.Get(int64(id))

	if err != nil {
		log.Printf("Erro ao pegar todo: %v", id)
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(todo)
}
