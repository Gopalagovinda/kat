package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

// infoHandler обрабатывает запросы и возвращает использованный HTTP-метод
func infoHandler(w http.ResponseWriter, r *http.Request) {
	v := fmt.Sprint(rand.IntN(6) + 1)
	w.Write([]byte(v))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/random", infoHandler)
	server := http.Server{
		Addr:    ":8083", // Обязательно двоеточие перед номером порта
		Handler: mux,     // Поле Handler должно быть с заглавной буквы
	}
	log.Printf("Сервер запущен и слушает порт %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
