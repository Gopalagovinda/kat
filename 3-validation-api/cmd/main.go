package main

import (
	"API/demo/configs"
	auth2 "API/demo/internal/auth"
	"log"
	"net/http"
)

func main() {
	// 1. Получаем конфигурацию (значение)
	conf := configs.GetConfig()

	// 2. Создаем маршрутизатор
	mux := http.NewServeMux()

	// 3. Передаем указатель на конфигурацию в конструктор хендлера
	auth2.NewAuthHandler(mux, &conf)

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
