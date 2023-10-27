package main

import (
	"fmt"
	"net/http"
	"time"
)

func StartServer(Addr int) error {

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", Addr),
		Handler:           HandlerRoutes(),
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	fmt.Printf("API starting at PORT %d \n", Addr)
	return srv.ListenAndServe()
}
