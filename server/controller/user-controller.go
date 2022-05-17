package controller

import (
	"example/hello/entity"
	"example/hello/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAllUsers() []entity.User
	SaveUser(ctx *gin.Context) entity.User
}

type controllerUser struct {
	service service.UserService
}

func NewUser(service service.UserService) UserController  {
	return &controllerUser {
		service: service,
	}
}

func (c *controllerUser) FindAllUsers() []entity.User {
	return c.service.FindAllUsers()
}

func (c *controllerUser) SaveUser(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.SaveUser(user)

	return user
}  