package usecase

import (
	"shiftboard/src/entity"

	"github.com/google/uuid"
)

type SpotInteractor struct {
	repository SpotRepositoryInterface
}

func NewSpotInteractor(spotRepository SpotRepositoryInterface) *SpotInteractor {
	return &SpotInteractor{
		repository: spotRepository,
	}
}

type SpotRepositoryInterface interface {
	Get(table string, user string, startWork string) (entity.TSpot, error)
	Post(table string, user string, spotId string, name string, salaly int, cutOffDay string, payDay string) (string, error)
}

func (interactor *SpotInteractor) DetailSpot(table string, user string, startWork string, spotId string) (entity.TSpot, error) {
	data, err := interactor.repository.Get(table, user, startWork+"_"+spotId)
	return data, err
}

func (interactor *SpotInteractor) CreateSpot(table string, user string, name string, salaly int, cutOffDay string, payDay string) (string, error) {
	spotId := uuid.New().String()
	data, err := interactor.repository.Post(table, user, spotId, name, salaly, cutOffDay, payDay)
	return data, err
}
