package main

import (
	"bougette-backend/cmd/api/handlers"
	"bougette-backend/common"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {
	e := echo.New()
	errenv := godotenv.Load()
	db, err := common.NewMysql()
	if err != nil {
		log.Fatal(err.Error())
	}
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	h := handlers.Handler{
		DB: db,
	}

	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}
	fmt.Println(app)
	port := os.Getenv("APP_PORT")
	appAdress := fmt.Sprintf("localhost:%s", port)
	e.Logger.Fatal(e.Start(appAdress))
}
