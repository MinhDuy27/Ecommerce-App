package service

import (
	"fmt"
	"go-app/domain"
	"go-app/internal/dto"
	"go-app/internal/repository"
)

type UserService struct{
	Repo repository.UserRepository
}

func (u *UserService) SignUp(input dto.SignUpdto) (string ,error) {
	usr := domain.User{
		Email: input.Email,
		Phone: input.Phone,
		Password: input.Password,
	}
	value ,err := u.Repo.CreateUser(usr)
	if err != nil{
		return "create faile",err
	}
	User_info:= fmt.Sprintf("%v,%v.%v",value.UserType,value.ID,value.Email)
	return User_info,nil
}

func (u *UserService) Login(input dto.Logindto) (string ,error) {
	value,err:= u.Repo.FindUserByEmail(input.Email)
	if err != nil{
		return "login failed",err
	}
	if value.Password != input.Password{
		return "incorrect Email or Password",nil
	}
	return value.Email,nil
}