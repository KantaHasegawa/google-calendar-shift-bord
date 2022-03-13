package usecase

import "shiftboard/src/entity"

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
}

func (interactor *SpotInteractor) DetailSpot(table string, user string, startWork string, spotId string) (entity.TSpot, error){
	data, err := interactor.repository.Get(table, user, startWork + "_" + spotId)
	return data, err
}
