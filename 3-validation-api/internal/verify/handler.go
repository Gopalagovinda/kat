package verify

import (
	"API/demo/configs"
	"fmt"
	"net/http"
)

type EmailHandler struct {
	*configs.Config
}

func NewEmailHandler(router *http.ServeMux, config *configs.Config) {
	handler := &EmailHandler{config}
	router.HandleFunc("POST /send", handler.send)
	router.HandleFunc("GET /verify/{hash}", handler.verify)
}

func (handler *EmailHandler) send(w http.ResponseWriter, r *http.Request) {

}

func (handler *EmailHandler) verify(w http.ResponseWriter, r *http.Request) {
	value := r.PathValue("hash")
	fmt.Println(value)
}
