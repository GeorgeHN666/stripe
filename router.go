package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandlerRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/express_create", CreateNewExpressAccount)
	mux.Get("/express_get", GetExpressConnectAccount)
	mux.Get("/del_acc", DelAccount)

	return mux
}
