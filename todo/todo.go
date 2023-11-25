package todo

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type HandlerFunc func(c *gin.Context) error

func List(c *gin.Context) {
	c.JSON(http.StatusOK, []Todo{})

}
