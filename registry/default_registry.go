package registry

import (
	"context"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/trangmaiq/kgs/internal/handler"
	"github.com/trangmaiq/kgs/persistence/sql"
)

var _ Registry = new(DefaultRegistry)

type DefaultRegistry struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup

	persister *sql.Persister
}

func New(engine *gin.Engine, db *gorm.DB) (Registry, error) {
	var (
		r   = new(DefaultRegistry)
		err = r.Init(context.TODO(), engine, db)
	)

	return r, err
}

func (r *DefaultRegistry) Init(_ context.Context, engine *gin.Engine, db *gorm.DB) error {
	r.engine = engine
	r.routerGroup = engine.Group("/keys")
	r.Handler().RegisterRoutes()

	r.persister = sql.NewPersister(db)

	return nil
}

func (r *DefaultRegistry) Handler() *handler.Handler {
	return handler.New(r)
}

func (r *DefaultRegistry) Routes() gin.IRoutes {
	return r.routerGroup
}

func (r *DefaultRegistry) Persister() handler.Persister {
	return r.persister
}
