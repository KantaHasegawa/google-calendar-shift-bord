package errorHandler

import (
	"fmt"
	"net/http"
)

func ControllerError(err error, w *http.ResponseWriter) {
	fmt.Println(err.Error())
	fmt.Fprint(*w, "sorry server error")
}
