package main

import (
	"log"

	"github.com/KKGo-Software-engineering/fun-exercise-api/helper"
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"os"

	"github.com/joho/godotenv"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("please create .env file with the following file .env.example")
	}

	dbConfig := postgres.Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	p, err := postgres.New(dbConfig)

	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Validator = helper.NewValidator()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(p)
	handler.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
