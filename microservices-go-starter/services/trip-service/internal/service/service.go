package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// интерфейс для хранилища поездок
type service struct {
	repo domain.TripRepository
}

// принимает репозиторий и возвращает указатель принятый этим репозиторием
func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

//ctx context.Context — контекст для контроля запроса.
//fare *domain.RideFareModel — модель тарифа поездки.

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{ // создает поездку
		ID:       primitive.NewObjectID(), //генерирует новый уникальный идентификатор
		UserID:   fare.UserID,             //устанавливает поле из переданного тарифа
		Status:   "pending",               //статус ожидание
		RideFare: fare,                    //задает тариф
	}

	return s.repo.CreateTrip(ctx, t) //возвращает поездку
}
