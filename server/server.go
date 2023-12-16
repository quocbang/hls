package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/quocbang/hls/http/handlers"
)

func Run() {
	e := echo.New()

	// init log

	// read config

	// middleware
	e.Use(middleware.CORS())

	// route
	handlers.NewHandler(e)

	// start
	log.Fatal(e.Start(":8000"))
}
