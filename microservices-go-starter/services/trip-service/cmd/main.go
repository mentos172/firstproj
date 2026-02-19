package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository() //создание репозитория который хранит данные

	svc := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID: "42",
	}

	t, err := svc.CreateTrip(ctx, fare) //вызов метода создания поездки
	if err != nil {
		log.Println(err)
	}

	log.Println(t) //выводим созданную поездку

	// keep the program running for now бесконечный мотоцикл
	for {
		time.Sleep(time.Second)
	}
}
