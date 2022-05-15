package registry

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Registry interface {
	Init(_ context.Context, engine *gin.Engine, db *gorm.DB) error
}
