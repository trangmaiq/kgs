package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	GetKeyResponse struct {
		Key string `json:"key"`
	}
	Persister interface {
		UseKey() (string, error)
	}
	PersistenceProvider interface {
		Persister() Persister
	}
	handlerDependencies interface {
		Routes() gin.IRoutes
		PersistenceProvider
	}
	Handler struct {
		hd handlerDependencies
	}
)

func New(hd handlerDependencies) *Handler {
	return &Handler{hd: hd}
}

func (h *Handler) RegisterRoutes() {
	h.hd.Routes().PUT("/use", h.useKey)
}

func (h *Handler) useKey(c *gin.Context) {
	key, err := h.hd.Persister().UseKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code": "internal_error",
			"message":    "get key failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"key": key,
	})
}
