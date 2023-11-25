package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type handler struct {
	message string
	id      string
}

func main() {
	r := gin.Default()
	h := handler{message: "pong", id: "id"}

	r.GET("/ping", h.ping)
	r.GET("/transfer/:id", h.transfer)

	err := r.Run()
	if err != nil {
		slog.Error(err.Error())
		return // exit
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (h handler) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": h.message,
	})
}

func (h handler) transfer(c *gin.Context) {
	id := c.Param(h.id)
	slog.Info("parsing...", slog.String("id", id))
	time.Sleep(time.Millisecond * 200)
	slog.Info("validating...", slog.String("id", id))
	time.Sleep(time.Millisecond * 100)
	slog.Info("staging...", slog.String("id", id))
	time.Sleep(time.Millisecond * 200)
	slog.Info("transaction starting...", slog.String("id", id))
	time.Sleep(time.Millisecond * 300)
	slog.Info("drawing...", slog.String("id", id))
	time.Sleep(time.Millisecond * 400)
	slog.Info("depositing...", slog.String("id", id))
	time.Sleep(time.Millisecond * 400)
	slog.Info("transaction ending...", slog.String("id", id))
	time.Sleep(time.Millisecond * 100)
	slog.Info("responding...", slog.String("id", id))
	time.Sleep(time.Millisecond * 100)
	slog.Info("finish", slog.String("id", id))
	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
