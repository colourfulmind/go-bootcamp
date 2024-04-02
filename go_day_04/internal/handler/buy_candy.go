package handler

import (
	"cmd/app/app.go/internal/event"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Input struct {
	Type  string `json:"candyType"`
	Count int    `json:"candyCount"`
	Money int    `json:"money"`
}

func (h *Handler) BuyCandyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var input Input
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				event.InputError(w, r, http.StatusBadRequest, err.Error())
				return
			}
			num, err := h.Candy.BuyCandy(input.Type, input.Money, input.Count)
			if err != nil {
				if errors.Is(err, event.WrongInput) || errors.Is(err, event.WrongType) {
					event.InputError(w, r, http.StatusBadRequest, err.Error())
				} else if errors.Is(err, event.NotEnoughMoney) {
					event.InputError(w, r, http.StatusPaymentRequired, fmt.Sprintf(err.Error(), num))
				}
			} else {
				event.InputSuccess(w, r, http.StatusCreated, num)
			}
		} else {
			event.RequestError(w, errors.New(fmt.Sprintf("method [%v] is not allowed", r.Method)).Error())
		}
	}
}
