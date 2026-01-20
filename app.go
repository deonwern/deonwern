package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"MementoMori11723/config"
	"MementoMori11723/server"
)

func main() {
	stopServer := startServer(&http.Server{
		Addr:    config.Defaults(),
		Handler: server.Mux(),
	})
	stopServer()
}

/*
	The startServer is a function that will not only start the server but also retun the stop functionality.
	The core idea is to treat it as a module, this allows us to keep it simple and only change it when needed,
	the function does the following:

	1. Listen to Ctrl + c signal or cancel signal & also creating errorHandler.
	2. Start the http server in a separate goroutine.
	3. A stop function is created to not only block the main function but also stop on cancel.

	Finally, we return the stop function.
*/

func startServer(serverConfig *http.Server) func() {
	stopSig := make(chan os.Signal, 1)
	signal.Notify(stopSig, os.Interrupt, syscall.SIGTERM)

	errorHandler := func(err error) {
		slog.Error(err.Error())
		os.Exit(1)
	}

	go func() {
		slog.Info("Starting Server", "url", serverConfig.Addr)
		if err := serverConfig.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errorHandler(err)
		}
	}()

	stopFunc := func() {
		<-stopSig
		slog.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := serverConfig.Shutdown(ctx); err != nil {
			errorHandler(err)
		}
	}

	return stopFunc
}
