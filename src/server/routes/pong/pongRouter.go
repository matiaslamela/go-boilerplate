package pongRouter

import (
	"github.com/gin-gonic/gin"
	pong "github.com/matiaslamela/go-boilerplate/src/server/controllers/pong"
)

type PongRoutes struct {
	Router *gin.RouterGroup
}

func (s *PongRoutes) Register() { //controlador?
	api := s.Router.Group("/pong")
	controller := pong.Controller{}
	api.GET("/", controller.Pong)
}
