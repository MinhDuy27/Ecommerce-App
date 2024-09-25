package service

import "go-app/internal/dto"

type UserService struct{}

func (u *UserService) SignUp(input dto.SignUpdto) (string ,error) {
	return "User Created",nil
}

func (u *UserService) Login(input dto.Logindto) (string ,error) {
	return "Login succes",nil
}