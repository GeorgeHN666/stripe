package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/account"
	"github.com/stripe/stripe-go/v72/accountlink"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/ephemeralkey"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/refund"
	"github.com/stripe/stripe-go/v72/setupintent"
)

const (
	SECRET_KEY = "sk_test_51O3wTiAml2Q20zJkhLIxBur3CAuMCZOdLlaHLK7zcXckKiYNmOXcH9oM4dr2hbsG5El0YqVqoymiINCxcaU3O5Kt00SDKCix4t"
	PUB_KEY    = "pk_test_51O3wTiAml2Q20zJkuxWMb6NtjKwq1DcxVBO8g0lPFvfd3vqMGuvQjU3aqGJbqvvatgl81TLDHCt62bqvEw8LjZz300UMlW0SJe"
	ELEVNT_FEE = 0.3
)

func InsertExpressAccount(Email string) (*stripe.Account, string, error) {
	stripe.Key = SECRET_KEY

	params := &stripe.AccountParams{
		Country: stripe.String("MX"),
		Type:    stripe.String(string(stripe.AccountTypeExpress)),
		Email:   &Email,

		Settings: &stripe.AccountSettingsParams{
			Payouts: &stripe.AccountSettingsPayoutsParams{
				DebitNegativeBalances: stripe.Bool(true),
				Schedule: &stripe.PayoutScheduleParams{
					DelayDays: stripe.Int64(1),
				},
			},
		},
	}

	account, err := account.New(params)
	if err != nil {
		return nil, "", err
	}
	fmt.Println("Account ==> ", account.ID)

	link, err := GenerateOnboardingLink(account.ID)
	if err != nil {
		return nil, "", err
	}

	fmt.Println("link", link)

	return account, link.URL, nil
}

func GenerateOnboardingLink(ID string) (*stripe.AccountLink, error) {

	stripe.Key = SECRET_KEY

	linkParam := &stripe.AccountLinkParams{
		Account:    stripe.String(ID),
		RefreshURL: stripe.String("https://app/refresh-onboarding"),
		ReturnURL:  stripe.String("https://test.zkaia.com"),
		Type:       stripe.String("account_onboarding"),
	}

	link, err := accountlink.New(linkParam)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func GetExpressAccount(ACC_ID string) (*stripe.Account, error) {

	stripe.Key = SECRET_KEY

	params := &stripe.AccountParams{}

	acc, err := account.GetByID(ACC_ID, params)
	if err != nil {
		return acc, err
	}

	return acc, nil
}

func DeleteAcc(ID string) error {

	stripe.Key = SECRET_KEY

	params := &stripe.AccountParams{}

	_, err := account.Del(ID, params)
	if err != nil {
		return err
	}
	return nil
}

func SimplePaymentIntent(amount int) (*stripe.PaymentIntent, error) {

	stripe.Key = SECRET_KEY

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(string(stripe.CurrencyMXN)),
		Customer: stripe.String("cus_Ow74WiNu6ozwJL"),
		Confirm:  stripe.Bool(true),
	}

	return paymentintent.New(params)
}

func CreatePaymenIntentWithFee(Amount int64, ACCID, CusID string) (*stripe.PaymentIntent, error) {

	stripe.Key = SECRET_KEY

	AmountLessFee := float64(Amount) * ELEVNT_FEE

	params := &stripe.PaymentIntentParams{
		Amount:               stripe.Int64(Amount),
		Currency:             stripe.String(string(stripe.CurrencyMXN)),
		Customer:             stripe.String(CusID),
		ApplicationFeeAmount: stripe.Int64(int64(AmountLessFee)),
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(ACCID),
		},
		OnBehalfOf: stripe.String(ACCID),
		Confirm:    stripe.Bool(true),
	}

	res, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreatePaymenIntentWithoutFee() {

}

func CreateStripePaymentSubscription(Email string) (*stripe.Customer, *stripe.EphemeralKey, *stripe.SetupIntent, error) {

	stripe.Key = SECRET_KEY

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(Email),
	}

	customer, err := customer.New(customerParams)
	if err != nil {
		return nil, nil, nil, err
	}

	ephiparams := &stripe.EphemeralKeyParams{
		Customer:      stripe.String(customer.ID),
		StripeVersion: stripe.String("2023-10-16"),
	}

	ephe, err := ephemeralkey.New(ephiparams)
	if err != nil {
		return nil, nil, nil, err
	}

	types := []*string{stripe.String("card")}

	param := &stripe.SetupIntentParams{
		Customer:           stripe.String(customer.ID),
		PaymentMethodTypes: types,
	}

	setupIntent, err := setupintent.New(param)
	if err != nil {
		return nil, nil, nil, err
	}

	return customer, ephe, setupIntent, nil
}

func CreateStripePaymentSubscriptionWithAccount(CUSTID string) (*stripe.EphemeralKey, *stripe.SetupIntent, error) {

	stripe.Key = SECRET_KEY

	ephiparams := &stripe.EphemeralKeyParams{
		Customer:      stripe.String(CUSTID),
		StripeVersion: stripe.String("2023-10-16"),
	}

	ephe, err := ephemeralkey.New(ephiparams)
	if err != nil {
		return nil, nil, err
	}

	types := []*string{stripe.String("card")}

	param := &stripe.SetupIntentParams{
		Customer:           stripe.String(CUSTID),
		PaymentMethodTypes: types,
	}

	setupIntent, err := setupintent.New(param)
	if err != nil {
		return nil, nil, err
	}

	return ephe, setupIntent, nil
}

func CreateRefundIntent(piID string) (*stripe.Refund, error) {

	stripe.Key = SECRET_KEY

	refundParams := &stripe.RefundParams{
		PaymentIntent: stripe.String(piID),
	}

	return refund.New(refundParams)

}
