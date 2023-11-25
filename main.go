package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"todoapi/todo"
)

type handler struct {
	message string
	id      string
}

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	h := handler{message: "pong", id: "id"}

	r.GET("/ping", h.ping)
	r.GET("/transfer/:id", h.transfer)
	r.GET("/todos", todo.List)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func (h handler) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": h.message,
	})
}

func (h handler) transfer(c *gin.Context) {
	id := c.Param("id")
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
		"message": h.message,
	})
}
