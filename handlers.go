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

func GetExpressConnectAccount(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error             bool            `json:"error"`
		Message           string          `json:"message"`
		Account           *stripe.Account `json:"account,omitempty"`
		InfoCompleted     bool            `json:"Info_completed,omitempty"`
		CanAcceptPayments bool            `json:"can_accept_payments,omitempty"`
		ID                string          `json:"id,omitempty"`
		OnBoardingURL     string          `json:"onboarding_url,omitempty"`
	}

	acc, err := GetExpressAccount(r.URL.Query().Get("acc"))
	if err != nil {
		fmt.Println("Huston we have a problem")
		res.Error = true
		res.Message = err.Error()
	}

	// if !acc.DetailsSubmitted {
	// 	fmt.Println("inside here")
	// 	link, err := GenerateOnboardingLink(acc.ID)
	// 	if err != nil {
	// 		res.Error = true
	// 		res.Message = err.Error()
	// 	} else {
	// 		res.OnBoardingURL = link.URL
	// 	}
	// }

	res.Error = false
	res.Message = "Account successfuly created"
	res.Account = acc
	res.InfoCompleted = acc.DetailsSubmitted
	res.CanAcceptPayments = acc.ChargesEnabled
	res.ID = acc.ID

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
