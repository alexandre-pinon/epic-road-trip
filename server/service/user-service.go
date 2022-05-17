package service

import (
	"example/hello/entity"
)

type UserService interface {
	SaveUser(entity.User) entity.User
	FindAllUsers() []entity.User
}

type userService struct {
	users []entity.User
}

func NewUser() UserService {
	return &userService{}
}

 func (service *userService) SaveUser(user entity.User) entity.User  {
	 service.users = append(service.users, user)
	 return user 
 }

 func (service *userService) FindAllUsers() []entity.User  {
	 return service.users
}