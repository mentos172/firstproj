package main

import (
	"log"      //создание логов
	"net/http" // работа с созданием сервера

	"ride-sharing/shared/env" //модуль в котором переменные
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081") // читаем переменну, и если ее нет то случаем порт
)

func main() {
	log.Println("Starting API Gateway")
	//запускаем сервер и на все запросы отвечаем хэло
	mux := http.NewServeMux()                                //вместо хттп встроенный мультиплексер
	mux.HandleFunc(" POST /trip/preview", handleTripPreview) // func(w http.ResponseWriter, r *http.Request) убрали это с кода и создали файл в это директории http
	//{
	//	w.WriteHeader(http.StatusOK)
	//	w.Write([]byte("Hello from API Gateway"))
	//})

	// http.ListenAndServe(httpAddr, nil) вариант до использования мультиплексера
	server := &http.Server{ //создание сервера с помощью мультиплексера
		Addr:    httpAddr,
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil { //если будет проблема с портом то перегрузится а так мы слушаем сервер
		log.Printf("HTTP server error: %v", err)
	}
}
