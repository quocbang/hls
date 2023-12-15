package server

import (
	"io"
	"log"
	"net/http"
	"os"

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

func RunMux() {
	http.HandleFunc("/watch", watch)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func watch(w http.ResponseWriter, r *http.Request) {
	videoFile, err := os.Open("./assets/video/den/mo_denvau.m3u8")
	if err != nil {
		http.Error(w, "Video file not found", http.StatusInternalServerError)
		return
	}
	defer videoFile.Close()

	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	io.Copy(w, videoFile)
	w.Write([]byte(`abc`))
}
