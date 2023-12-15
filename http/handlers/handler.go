package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/quocbang/hls/http/handlers/video"
)

func NewHandler(e *echo.Echo) error {
	api := e.Group("/api")

	video.Init(api)
	return nil
}
