package video

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type service struct {
}

func NewVideoService() Video {
	return &service{}
}

func (s *service) Watch(c echo.Context) error {
	c.Response().Header().Set("Content-type", "application/vnd.apple.mpegurl")
	return c.File("./assets/video/master.m3u8")
}

type PlayVideoRequest struct {
	Size string `json:"size" query:"size"`
	File string `json:"file" query:"file"`
}

func (s *service) Play(c echo.Context) error {
	var (
		req PlayVideoRequest
	)
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}
	log.Printf("start play size: %s file: %s", req.Size, req.File)
	return c.File(fmt.Sprintf("./assets/video/%s/%s", req.Size, req.File))
}
