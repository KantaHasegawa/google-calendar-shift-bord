package usecase

type ShiftInteractor struct {
	repository ShiftRepositoryInterface
}

type ShiftRepositoryInterface interface{

}

func NewShiftInteractor(repository ShiftRepositoryInterface)(*ShiftInteractor){
	return &ShiftInteractor{repository: repository}
}
