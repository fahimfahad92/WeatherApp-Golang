package main

import (
	"WeatherApp/application"
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
)

func main() {

	log.Println("Weather App")

	app := application.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	configureViper()

	err := app.Start(ctx)
	if err != nil {
		log.Println("failed to start weather app:", err)
	}
}

func configureViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
}
