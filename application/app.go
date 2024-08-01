package application

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{}
	loadRoutes(app)
	return app
}

func (a *App) Start(ctx context.Context) error {

	port := fmt.Sprintf(":%s", viper.GetString("port"))
	log.Printf("Server port is: %s\n", port)

	server := &http.Server{
		Addr:    port,
		Handler: a.router,
	}

	ch := make(chan error, 1)

	var err error
	go func() {
		err, ch = startServer(err, server, ch)
	}()

	return processChannelResponse(ctx, err, ch, server)
}

func processChannelResponse(ctx context.Context, err error, ch chan error, server *http.Server) error {
	select {
	case err = <-ch:
		log.Println("Received error:", err)
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		log.Println("Shutting down server")
		return server.Shutdown(timeout)
	}
}

func startServer(err error, server *http.Server, ch chan error) (error, chan error) {
	log.Println("Server starting")
	err = server.ListenAndServe()
	if err != nil {
		ch <- fmt.Errorf("failed to start server: %w", err)
	}

	close(ch)
	return err, ch
}
