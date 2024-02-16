package handlers

import (
	"encoding/json"
	"fmt"
	"gomod/models"
	"net/http"
)

func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	responseMessage := models.Message{"Welcome to Rest API Project"}
	data, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Println("Hata:", err.Error())
	}
	fmt.Fprintf(w, string(data))
}
