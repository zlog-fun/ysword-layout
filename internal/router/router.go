package router

import (
	"ysword_layout/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Routers struct {
	Blog *service.BlogService
}

// Register
func (r *Routers) Register(engine *gin.Engine) {
	// TODO according to appname, use router
	engine.GET("/ping", func(context *gin.Context) {
		context.String(200, "v0.0.1")
	})
	engine.GET("/", r.Blog.Index)
}

// manaual reagister service
func New(b *service.BlogService) func(*gin.Engine) {
	r := &Routers{
		Blog: b,
	}
	return func(engine *gin.Engine) {
		r.Register(engine)
	}
}

var ProviderSet = wire.NewSet(New)
