package routes

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
