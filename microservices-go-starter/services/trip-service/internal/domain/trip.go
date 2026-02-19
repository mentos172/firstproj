package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct { // это мы храним в базе данных
	ID       primitive.ObjectID //уникальный идентификатор
	UserID   string             //идентификатор пользователя которому пренодлежит поездка
	Status   string             //статус поездки мы возвращаем пендинг
	RideFare *RideFareModel     //указатель на структура которая описывает тариф цены расчет
}

// описываем слой хранинея данных
// сохраняет поездку и возращает сохраненную
type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

// описываем логику работы с пездками
// создается поездка на основе тарифа
type TripService interface {
	CreateTrip(ctx context.Context, trip *RideFareModel) (*TripModel, error)
}
