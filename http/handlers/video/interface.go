package video

import "github.com/labstack/echo/v4"

type Video interface {
	Watch(echo.Context) error
	Play(echo.Context) error
}
