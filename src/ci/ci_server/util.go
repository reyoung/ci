package ci_server

import (
	"log"
	"fmt"
	"net/http"
)

func CheckNoErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}


func makeSafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				http.Error(w, fmt.Sprint(e), http.StatusInternalServerError)
			}
		}()
		fn(w, r)
	}
}