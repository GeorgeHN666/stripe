package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandlerRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/express_create", CreateNewExpressAccount)
	mux.Get("/express_get", ContinueOnboarding)
	mux.Get("/del_acc", DelAccount)
	mux.Get("/setup_payment", GetSetupIntent)
	mux.Get("/buy", BuyWithFee)
	mux.Get("/simple", SimplePI)
	mux.Get("/refund", CreateRefundIntentEP)

	// operational
	mux.Get("/items", GetItems)

	return mux
}
