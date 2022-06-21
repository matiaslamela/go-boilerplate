package httpUtils

import "github.com/gin-gonic/gin"

func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
type HttpJson struct {
	Data any `json:"data"`
}
type HttpJsonList struct {
	Data  any `json:"data"`
	Total int `json:"total" example:"8"`
	Page  int `page:"total" example:"1"`
}

func SendJsonStruct[T any](ctx *gin.Context, status int, data T) {
	ctx.JSON(status, HttpJson{Data: data})
}

func SendJsonList[T any](ctx *gin.Context, status int, data []T, page int, total int) {
	ctx.JSON(status, HttpJsonList{
		Data:  data,
		Total: total,
		Page:  page,
	})
}
