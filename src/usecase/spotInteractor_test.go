package usecase

import (
	"testing"
)

func TestValidateName(t *testing.T) {
	arg_20 := "01234567890123456789"
	arg_21 := "012345678901234567890"
	err1 := nameValidate(arg_20)
	if err1 != nil {
		t.Errorf("faild name validate test")
	}
	err2 := nameValidate(arg_21)
	if err2 == nil {
		t.Errorf("faild name validate test")
	}
}
