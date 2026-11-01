package auth

import (
	"API/demo/configs"
	"encoding/json"
	"fmt"
	"net/http"
)

type Authhandler struct {
	*configs.Config // Внедрение зависимости через встраивание (embedding)
}

// Принимаем только конфигурацию, а не весь хендлер
func NewAuthHandler(router *http.ServeMux, cfg *configs.Config) {
	// Создаем хендлер один раз
	handler := &Authhandler{Config: cfg}

	// Регистрируем методы напрямую (работает в Go 1.22+)
	router.HandleFunc("POST /auth/register", handler.register)
	router.HandleFunc("GET /auth/login", handler.login)
}

// Сигнатура сразу соответствует тому, что ожидает HandleFunc
func (h *Authhandler) login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("your token: %s\n", h.Auth.Secret)
	w.Header().Set("Content-Type", "application/json")
	data := LoginResponse{
		Token: h.Auth.Secret,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *Authhandler) register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Register successful"))
}
