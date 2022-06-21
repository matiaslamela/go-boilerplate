package user

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/matiaslamela/go-boilerplate/src/database/user"
	"github.com/matiaslamela/go-boilerplate/src/utils/httpUtils"
)

type Controller struct {
	Repository userModel.Repository
}

// ShowAccount godoc
// @Summary      GetUsers
// @Description  Gets all the users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  httpUtils.HttpJsonList{data=[]userModel.User}
// @Router       /users [get]
func (uc *Controller) GetUsers(c *gin.Context) {
	users, err := uc.Repository.GetAll(c)
	if err != nil {
		httpUtils.NewError(c, http.StatusInternalServerError, err)
		return
	}
	httpUtils.SendJsonList(c, http.StatusOK, *users, 1, 10)
}

func (uc *Controller) CreateUser(c *gin.Context) {
	var u userModel.User
	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	user, createErr := uc.Repository.Create(c, u)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
