package service

import (
	"example/hello/entity"
)

type ActivityService interface {
	SaveActivity(entity.Activity) entity.Activity
	FindAllActivities() []entity.Activity
}

type activityService struct {
	activities []entity.Activity
}

 func NewActivity() ActivityService {
	 return &activityService{}
 }

 func (service *activityService) SaveActivity(activity entity.Activity) entity.Activity  {
	 service.activities = append(service.activities, activity)
	 return activity 
 }

 func (service *activityService) FindAllActivities() []entity.Activity  {
	 return service.activities
}