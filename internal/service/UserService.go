package service

import (
	"go-app/domain"
	"go-app/internal/dto"
	"go-app/internal/helper"
	"go-app/internal/repository"
	"log"
	"strconv"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (u *UserService) SignUp(input dto.SignUpdto) (string, error) {
	hashedPassword, err := u.Auth.HashPassword(input.Password)
	if err != nil {
		return "",err
	}
	usr := domain.User{
		Email:    input.Email,
		Phone:    input.Phone,
		Password: hashedPassword,
	}
	value, err := u.Repo.CreateUser(usr)
	if err != nil {
		return "", err
	}
	token, err := u.Auth.GenerateToken(value.ID, value.Email, value.UserType)
	if err != nil {
		return "",err
	}
	return token, nil
}

func (u *UserService) Login(input dto.Logindto) (string, error) {
	value, err := u.Repo.FindUserByEmail(input.Email)
	if err != nil {
		return "", err
	}
	err = u.Auth.VerifyPassword(value.Password,input.Password)
	if err != nil {
		return "", err
	}
	token, err := u.Auth.GenerateToken(value.ID, value.Email, value.UserType)
	if err != nil {
		return "",err
	}
	return  token,nil
}
func (u *UserService) GetProfilesByID(idInput string) (domain.User, error) {
	idInt, err := strconv.ParseUint(idInput,10,32)
	if err != nil {
		return domain.User{}, err
	}
	value, error := u.Repo.FindUserById(uint(idInt))
	if error != nil {
		return domain.User{}, error
	}
	return value, nil

}
func (u *UserService) GetProfilesByEmail(email string) (domain.User, error) {
	value, error := u.Repo.FindUserByEmail(email)
	if error != nil {
		return domain.User{}, error
	}
	return value, nil

}
func (u *UserService) UpdateUser(idInput string, user domain.User) error {
	idInt, err := strconv.ParseUint(idInput,10,32)
	if err != nil {
		return err
	}
	value, err := u.Repo.UpdateUser(uint(idInt), user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) CreateProfile(id string, p dto.CreateProfiledto) error {
	idInt, err := strconv.ParseUint(id,10,32)
	if err != nil {
		return err
	}
	user := domain.User{
		Phone:     p.Phone,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
	value, err := u.Repo.UpdateUser(uint(idInt), user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) BecomeSeller(idInput string) error {
	idInt, err := strconv.ParseUint(idInput,10,32)
	if err != nil {
		return err
	}
	user := domain.User{
		UserType: "seller",
	}
	value, err := u.Repo.UpdateUser(uint(idInt), user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) RevokeSeller(idInput string) error {
	idInt, err := strconv.ParseUint(idInput,10,32)
	if err != nil {
		return err
	}
	user := domain.User{
		UserType: "buyer",
	}
	value, err := u.Repo.UpdateUser(uint(idInt), user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
