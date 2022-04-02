package usecase

import "shiftboard/src/entity"

type ShiftInteractor struct {
	repository ShiftRepositoryInterface
}

type ShiftRepositoryInterface interface{
	Get(table string, user string, year string, month string)([]entity.TShift, error)
	Post(table string, user string, startWork string, finishWork string, spotId string)(error)
	Delete(table string, user string, startWork string)(error)
}

func NewShiftInteractor(repository ShiftRepositoryInterface)(*ShiftInteractor){
	return &ShiftInteractor{repository: repository}
}

func (interactor *ShiftInteractor) IndexShift(table string, user string, year string, month string)([]entity.TShift, error){
	result, err := interactor.repository.Get(table, user, year, month)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (interactor *ShiftInteractor) CreateShift(table string, user string, startWork string, finishWork string, spotId string)(error){
	err := interactor.repository.Post(table, user, startWork, finishWork, spotId)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *ShiftInteractor) DeleteShift(table string, user string, startWork string)(error){
	err := interactor.repository.Delete(table, user, startWork)
	if err != nil {
		return err
	}
	return nil
}
