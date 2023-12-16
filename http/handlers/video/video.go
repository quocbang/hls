package video

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

type service struct {
}

func NewVideoService() Video {
	return &service{}
}

func (s *service) Watch(c echo.Context) error {
	log.Println("start quest video")
	f, err := os.Open("./assets/video/den/mo_denvau.m3u8")
	if err != nil {
		return c.JSON(500, fmt.Errorf("failed to open video, error: %v", err).Error())
	}
	defer f.Close()
	c.Response().Header().Set("Content-type", "application/vnd.apple.mpegurl")
	return c.File("./assets/video/den/mo_denvau.m3u8")
}

func (s *service) Play(c echo.Context) error {
	requestFile := c.Param("requestFile")
	log.Printf("start play file %s", requestFile)
	return c.File(fmt.Sprintf("./assets/video/den/%s", requestFile))
}
