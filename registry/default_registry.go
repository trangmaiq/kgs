package registry

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/trangmaiq/kgs/internal/handler"
)

var _ Registry = new(DefaultRegistry)

type DefaultRegistry struct {
	engine *gin.Engine

	urlRouterGroup *gin.RouterGroup
}

func New(engine *gin.Engine) (Registry, error) {
	var (
		r   = new(DefaultRegistry)
		err = r.Init(context.TODO(), engine)
	)

	return r, err
}

func (r *DefaultRegistry) Init(_ context.Context, engine *gin.Engine) error {
	r.engine = engine
	r.urlRouterGroup = engine.Group("/keys")

	r.Handler().RegisterRoutes()

	return nil
}

func (r *DefaultRegistry) Handler() *handler.Handler {
	return handler.New(r)
}

func (r *DefaultRegistry) Routes() gin.IRoutes {
	return r.engine
}

func (r *DefaultRegistry) Persister() handler.Persister {
	//TODO implement me
	panic("implement me")
}
