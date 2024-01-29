package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// PageError outputs an error if the page doesn't exist
func PageError(w http.ResponseWriter, err error, page int) {
	var text string
	if err != nil {
		text = strings.Trim(strings.Split(err.Error(), " ")[2], ":")
		text = strings.TrimLeft(text, "\"")
		text = strings.TrimRight(text, "\"")
	} else {
		text = strconv.Itoa(page)
	}
	log.Println(fmt.Sprintf("Invalid 'page' value: '%v'", text))
	http.Error(w, fmt.Sprintf("Invalid 'page' value: '%v'", text), http.StatusBadRequest)
}

// ResponseError outputs an error if one occurred while sending a response
func ResponseError(w http.ResponseWriter, err error) {
	text := fmt.Sprintf("[" + strconv.Itoa(http.StatusInternalServerError) + "] internal server error")
	log.Println(text)
	http.Error(w, text, http.StatusInternalServerError)
}

// RequestError outputs an error if a bad request is sent
func RequestError(w http.ResponseWriter, err error) {
	text := fmt.Sprintf("["+strconv.Itoa(http.StatusBadRequest)+"] %s", err.Error())
	log.Println(text)
	http.Error(w, text, http.StatusBadRequest)
}

// UnauthorizedError outputs an error if the user wasn't authorized
func UnauthorizedError(w http.ResponseWriter, err error) {
	text := fmt.Sprintf("["+strconv.Itoa(http.StatusUnauthorized)+"] %s", err.Error())
	log.Println(text)
	http.Error(w, text, http.StatusUnauthorized)
}
