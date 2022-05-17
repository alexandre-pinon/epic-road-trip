package controller

import (
	"example/hello/entity"
	"example/hello/service"
	"github.com/gin-gonic/gin"
)

type ActivityController interface {
	FindAllActivities() []entity.Activity
	SaveActivity(ctx *gin.Context) entity.Activity
}

type controllerActivity struct {
	service service.ActivityService
}

func NewActivity(service service.ActivityService) ActivityController  {
	return &controllerActivity {
		service: service,
	}
}

func (c *controllerActivity) FindAllActivities() []entity.Activity {
	return c.service.FindAllActivities()
}

func (c *controllerActivity) SaveActivity(ctx *gin.Context) entity.Activity {
	var activity entity.Activity
	ctx.BindJSON(&activity)
	c.service.SaveActivity(activity)

	return activity
}  