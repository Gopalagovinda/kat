package main

import (
	"log"
	"math/rand/v2"
	"net/http"
)

// infoHandler обрабатывает запросы и возвращает использованный HTTP-метод
func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string(rand.IntN(6))))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/random", infoHandler)
	server := http.Server{
		Addr:    ":8082", // Обязательно двоеточие перед номером порта
		Handler: mux,     // Поле Handler должно быть с заглавной буквы
	}
	log.Printf("Сервер запущен и слушает порт %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
