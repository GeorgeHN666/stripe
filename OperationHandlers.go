package main

import "net/http"

func InsertUserEP(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		User    *User  `json:"user"`
	}

	var User User

	err := ReadJSON(w, r, &User)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	UserData, err := StartDB().InsertUser(&User)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.User = UserData
	res.Error = false
	res.Message = "All good"
	WriteJSON(w, r, http.StatusOK, res)
}

func GetUserEP(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		User    *User  `json:"user"`
	}

	user, err := StartDB().GetUser(r.URL.Query().Get("id"))
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.User = user

	res.Error = false
	res.Message = "All good"
	WriteJSON(w, r, http.StatusOK, res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	var user User

	err := ReadJSON(w, r, user)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	err = StartDB().UpdateUser(&user, r.URL.Query().Get("id"))
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "All good"

	WriteJSON(w, r, http.StatusOK, res)
}

func InsertItemsEP(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	var Items []*Item

	err := ReadJSONTo(w, r, Items)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	err = StartDB().InsertItems(Items)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "All good"
	WriteJSON(w, r, http.StatusOK, res)

}

func GetItems(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool    `json:"error"`
		Message string  `json:"message"`
		Items   []*Item `json:"items"`
	}

	items, err := StartDB().GetItems()
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "All good"
	res.Items = items
	WriteJSON(w, r, http.StatusOK, res)
}

func CreatePaymentIntentEP(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	var PI PaymentIntent

	err := ReadJSON(w, r, &PI)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	err = StartDB().CreatePaymentIntent(&PI)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "All good"
	WriteJSON(w, r, http.StatusOK, res)

}

func GetPaymentIntentEP(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error         bool           `json:"error"`
		Message       string         `json:"message"`
		PaymentIntent *PaymentIntent `json:"payment_intent"`
	}

	paymentIntent, err := StartDB().GetPaymentIntent(r.URL.Query().Get("pi"))
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "All good"
	res.PaymentIntent = paymentIntent
	WriteJSON(w, r, http.StatusOK, res)

}

func CreateRefundEP(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	var refund Refunds

	err := ReadJSON(w, r, &refund)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	err = StartDB().CreateRefundIntent(&refund)
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "All good"
	WriteJSON(w, r, http.StatusOK, res)

}

func GetPaymentIntent(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error        bool     `json:"error"`
		Message      string   `json:"message"`
		RefundIntent *Refunds `json:"refund"`
	}

	refundIntent, err := StartDB().GetRefunIntent(r.URL.Query().Get("refund"))
	if err != nil {
		res.Error = true
		res.Message = err.Error()

		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "All good"
	res.RefundIntent = refundIntent
	WriteJSON(w, r, http.StatusOK, res)

}
