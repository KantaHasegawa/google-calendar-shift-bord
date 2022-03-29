package usecase

import (
	"errors"
	"shiftboard/src/entity"
	"strconv"
	"strings"

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
	//userが存在するか
	//nameの重複判定
	err := nameValidate(name)
	if(err != nil) {
		return "", err
	}
	err = cutOffDayValidate(cutOffDay)
	if(err != nil) {
		return "", err
	}
	err = payDayValidate(payDay)
	if(err != nil) {
		return "", err
	}
	err = daysValidate(cutOffDay, payDay)
	if(err != nil) {
		return "", err
	}
	spotId := uuid.New().String()
	data, err := interactor.repository.Post(table, user, spotId, name, salaly, cutOffDay, payDay)
	return data, err
}

func cutOffDayValidate(cutOffDay string) error {
	validateList := make([]string, 29)
	for i := range validateList {
		if i < 10 {
			validateList[i] = "0" + strconv.Itoa(i)
		}
		validateList[i] = strconv.Itoa(i)
	}
	validateList[0] = "EOM"
	for _, item := range validateList {
		if cutOffDay == item {
			return nil
		}
	}
	return errors.New("締切日が不正です")
}

func payDayValidate(payDay string) error {
	validateList := make([]string, 60)
	for i := 1; i < 29; i++ {
		if i < 10 {
			validateList = append(validateList, "CURRENT_"+"0"+strconv.Itoa(i), "NEXT_"+"0"+strconv.Itoa(i))
		}
		validateList = append(validateList, "CURRENT_"+strconv.Itoa(i), "NEXT_"+strconv.Itoa(i))
	}
	validateList = append(validateList, "CURRENT_"+ "EOM", "NEXT_"+ "EOM")
	for _, item := range validateList {
		if payDay == item {
			return nil
		}
	}
	return errors.New("給料日が不正です")
}

func daysValidate(cutOffDay string, payDay string) error {
	arr := strings.Split(payDay, "_")
	if arr[0] == "NEXT" {
		return nil
	}
	if cutOffDay > arr[1] {
		return errors.New("給料日と締切日が不正です")
	}
	return nil
}

func nameValidate(name string) error {
	if(len(name) == 0 || len(name) > 20){
		return errors.New("名前の長さが不正です")
	}
	return nil
}
