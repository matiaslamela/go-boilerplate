package userRouter

import (
	"github.com/gin-gonic/gin"
	database "github.com/matiaslamela/go-boilerplate/src/database"
	userRepository "github.com/matiaslamela/go-boilerplate/src/repositories/user"
	userController "github.com/matiaslamela/go-boilerplate/src/server/controllers/user"
)

type UserRoutes struct {
	Router *gin.RouterGroup
}

func (ur *UserRoutes) Register() {
	api := ur.Router.Group("/users")
	controller := userController.Controller{Repository: &userRepository.Repository{
		DB: database.Conn.DB,
	}}
	api.GET("/", controller.GetUsers)
	api.POST("/", controller.CreateUser)
}
