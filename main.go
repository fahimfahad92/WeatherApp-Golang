package main

import (
	"WeatherApp/application"
	"WeatherApp/config"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Weather App")

	app := application.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	config.ConfigureViper()

	err := app.Start(ctx)
	if err != nil {
		log.Println("failed to start weather app:", err)
	}
}
