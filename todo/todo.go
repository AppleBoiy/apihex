package todo

import (
	"database/sql"
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

type Handler struct {
	db *sql.DB
}

type HandlerFunc func(c *gin.Context) error

func NewHandle(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h Handler) NewTask(c *gin.Context) {
	var t Todo
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	_, err := h.db.Exec("INSERT INTO todos (title) VALUES (?)", t.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

}

func (h Handler) List(c *gin.Context) {
	c.JSON(http.StatusOK, []Todo{})
}
