package video

import (
	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	video := NewVideoService()

	g.GET("/watch", video.Watch)
}
