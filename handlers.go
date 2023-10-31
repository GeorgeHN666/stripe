package main

import (
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v72"
)

// Errors
const (
	USER_EXIST      = 500
	USER_DONT_EXIST = 200
)

func CreateNewExpressAccount(w http.ResponseWriter, r *http.Request) {
	// Create account if the user doesn't exist

	var res struct {
		Error   bool            `json:"error"`
		Message string          `json:"message"`
		Account *stripe.Account `json:"account"`
		Link    string          `json:"link"`
	}

	acc, link, err := InsertExpressAccount(r.URL.Query().Get("email"))
	if err != nil {
		res.Error = true
		res.Message = err.Error()
	}

	fmt.Println(acc)

	res.Error = false
	res.Message = "Account successfuly created"
	res.Account = acc
	res.Link = link

	WriteJSON(w, r, http.StatusOK, res)
}

func ContinueOnboarding(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error         bool   `json:"error"`
		Message       string `json:"message"`
		ID            string `json:"id"`
		OnBoardingURL string `json:"onboarding_url"`
	}

	link, err := GenerateOnboardingLink(r.URL.Query().Get("acc"))
	if err != nil {
		res.Error = true
		res.Message = err.Error()
	} else {
		res.OnBoardingURL = link.URL
	}

	res.Error = false
	res.Message = "Account successfuly created"

	WriteJSON(w, r, http.StatusOK, res)
}

func DelAccount(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	err := DeleteAcc(r.URL.Query().Get("acc"))
	if err != nil {
		res.Error = true
		res.Message = err.Error()
	}

	WriteJSON(w, r, http.StatusOK, res)

}

func GetSetupIntent(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Error          bool   `json:"error"`
		Message        string `json:"message"`
		EphemeralKey   string `json:"ephemeral_key"`
		SetupIntentKey string `json:"setup_intent_key"`
		Customer       string `json:"customer_id"`
	}

	customer, ephe, setup, err := CreateStripePaymentSubscription(r.URL.Query().Get("acc"))
	if err != nil {
		res.Error = true
		res.Message = err.Error()
	} else {
		res.EphemeralKey = ephe.ID
		res.SetupIntentKey = setup.ClientSecret
		res.Customer = customer.ID
	}

	WriteJSON(w, r, http.StatusOK, res)
}
