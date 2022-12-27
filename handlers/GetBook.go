package handlers

import (
	"encoding/json"
	"fmt"
	"graphql-go/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	if book.Id == 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}

}
