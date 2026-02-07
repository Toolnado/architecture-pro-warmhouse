package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Addr    string
	Handler http.Handler
}

func NewApp(addr string, h http.Handler) *App {
	return &App{
		Addr:    addr,
		Handler: h,
	}
}

func (app *App) Run() error {
	server := http.Server{
		Addr:    app.Addr,
		Handler: app.Handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("failed to start server: %s", err)
		}
	}()

	log.Println("Server starting on ", app.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	return server.Shutdown(context.Background())
}
