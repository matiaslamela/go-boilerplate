package pong

import (
	"github.com/gin-gonic/gin"
	"github.com/matiaslamela/go-boilerplate/src/utils/httpUtils"
)

type Controller struct{}

type Pong struct {
	Ping string `json:"ping" example:"pong"`
}

// ShowAccount godoc
// @Summary      Piiiiiiinggg
// @Description  Pooooooooong
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  httpUtils.HttpJson{data=Pong}  "pong"
// @Router       /pong [get]
func (*Controller) Pong(c *gin.Context) {
	httpUtils.SendJsonStruct(c, 200, Pong{Ping: "pong"})
}
