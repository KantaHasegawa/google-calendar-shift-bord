package entity

type TShift struct {
	User string `json:"User"`
	StartWork string `json:"StartWork"`
	FinishWork string `json:"FinishWork"`
	SpotId string `json:"SpotId"`
}

type ShiftInteractorInteface interface{
	CreateShift(table string, user string, startWork string, finishWork string, spotId string)(error)
}
