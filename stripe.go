package main

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/account"
	"github.com/stripe/stripe-go/v72/accountlink"
)

const (
	SECRET_KEY = "sk_test_51O3wTiAml2Q20zJkhLIxBur3CAuMCZOdLlaHLK7zcXckKiYNmOXcH9oM4dr2hbsG5El0YqVqoymiINCxcaU3O5Kt00SDKCix4t"
	PUB_KEY    = "pk_test_51O3wTiAml2Q20zJkuxWMb6NtjKwq1DcxVBO8g0lPFvfd3vqMGuvQjU3aqGJbqvvatgl81TLDHCt62bqvEw8LjZz300UMlW0SJe"
)

func InsertExpressAccount(Email string) (*stripe.Account, string, error) {

	params := &stripe.AccountParams{
		Country: stripe.String("MX"),
		Type:    stripe.String(string(stripe.AccountTypeExpress)),
		Email:   &Email,
	}

	account, err := account.New(params)
	if err != nil {
		return nil, "", err
	}

	link, err := GenerateOnboardingLink(account.ID)
	if err != nil {
		return nil, "", err
	}

	return account, link.URL, nil
}

func GenerateOnboardingLink(ID string) (*stripe.AccountLink, error) {
	linkParam := &stripe.AccountLinkParams{
		Account:    &ID,
		RefreshURL: stripe.String("payment-processing://refresh-onboarding"),
		ReturnURL:  stripe.String("payment-processing://stripe-return"),
		Type:       stripe.String("aacount_onboarding"),
	}

	link, err := accountlink.New(linkParam)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func GetExpressAccount(ACC_ID string) (*stripe.Account, error) {

	params := &stripe.AccountParams{}

	acc, err := account.GetByID(ACC_ID, params)
	if err != nil {
		return nil, err
	}

	return acc, nil
}