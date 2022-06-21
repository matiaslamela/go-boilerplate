package mainRouter

import (
	"github.com/gin-gonic/gin"
	pongRouter "github.com/matiaslamela/go-boilerplate/src/server/routes/pong"
	userRouter "github.com/matiaslamela/go-boilerplate/src/server/routes/user"
)

type Routes struct {
	App *gin.Engine
	api *gin.RouterGroup
}

type Router interface {
	Register()
}

func AddRouter(r Router) {
	r.Register()
}

func (s *Routes) Register() { //controlador?
	api := s.App.Group("/api/v1")
	AddRouter(&pongRouter.PongRoutes{Router: api})
	AddRouter(&userRouter.UserRoutes{Router: api})
}
