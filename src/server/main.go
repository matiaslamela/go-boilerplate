package router

import (
	"github.com/gin-gonic/gin"
	enviroment "github.com/matiaslamela/go-boilerplate/src/env"
	mainRouter "github.com/matiaslamela/go-boilerplate/src/server/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/matiaslamela/go-boilerplate/docs" // docs is generated by Swag CLI, you have to import it.
)

type Server struct {
	app *gin.Engine
}

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func (s *Server) Start() { //controlador?
	s.app = gin.New()

	docs.SwaggerInfo.Title = "Golang example"
	docs.SwaggerInfo.Description = "This is a sample server fo golang."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "localhost:3001"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	s.app.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := mainRouter.Routes{App: s.app}
	router.Register()
	s.app.Run(":" + enviroment.PORT)
}
