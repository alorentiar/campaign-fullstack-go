package handler

import (
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	//catch input user
	//mapping input dari user
	//struct dilakukan passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err!=nil{
		c.JSON(http.StatusBadRequest, nil)
	}

	user,err := h.userService.RegisterUser(input)
	if err!=nil {
		c.JSON(http.StatusBadRequest,nil)
	}

	c.JSON(http.StatusOK, user)
}