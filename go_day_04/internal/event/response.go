package event

// #include "cow.h"
import "C"

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"unsafe"
)

type ResponseSuccess struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}

func InputSuccess(w http.ResponseWriter, r *http.Request, Status int, change int) {
	log.Printf("[%s] %s - Response: %d", r.Method, r.URL.Path, Status)
	ty := C.CString("Thank you!")
	answer := C.ask_cow(ty)
	defer C.free(unsafe.Pointer(answer))
	resp, _ := json.Marshal(ResponseSuccess{
		Change: change,
		Thanks: C.GoString(answer),
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(Status)
	w.Write(resp)
}

var (
	WrongInput     = errors.New("wrong input")
	WrongType      = errors.New("wrong type of candy")
	NotEnoughMoney = errors.New("you need %d more money")
)

type ResponseError struct {
	Error string `json:"error"`
}

func InputError(w http.ResponseWriter, r *http.Request, Status int, err string) {
	log.Printf("[%s] %s - Error: %s", r.Method, r.URL.Path, err)
	resp, _ := json.Marshal(ResponseError{
		Error: err,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(Status)
	w.Write(resp)
}

func RequestError(w http.ResponseWriter, err string) {
	log.Println(err)
	resp, _ := json.Marshal(ResponseError{
		Error: err,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp)
}
