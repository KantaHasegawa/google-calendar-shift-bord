package usecase

type ShiftInteractor struct {
	repository ShiftRepositoryInterface
}

type ShiftRepositoryInterface interface{
	Post(table string, user string, startWork string, finishWork string, spotId string)(error)
}

func NewShiftInteractor(repository ShiftRepositoryInterface)(*ShiftInteractor){
	return &ShiftInteractor{repository: repository}
}

func (interactor *ShiftInteractor) CreateShift(table string, user string, startWork string, finishWork string, spotId string)(error){
	err := interactor.repository.Post(table, user, startWork, finishWork, spotId)
	if err != nil {
		return err
	}
	return nil
}
