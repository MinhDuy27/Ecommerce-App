package service

import (
	"fmt"
	"go-app/domain"
	"go-app/internal/dto"
	"go-app/internal/repository"
	"log"
	"strconv"
)

type UserService struct {
	Repo repository.UserRepository
}

func (u *UserService) SignUp(input dto.SignUpdto) (string, error) {
	usr := domain.User{
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	}
	value, err := u.Repo.CreateUser(usr)
	if err != nil {
		return "create faile", err
	}
	User_info := fmt.Sprintf("%v,%v.%v", value.UserType, value.ID, value.Email)
	return User_info, nil
}

func (u *UserService) Login(input dto.Logindto) (string, error) {
	value, err := u.Repo.FindUserByEmail(input.Email)
	if err != nil {
		return "login failed", err
	}
	if value.Password != input.Password {
		return "incorrect Email or Password", nil
	}
	return value.Email, nil
}
func (u *UserService) GetProfilesByID(idInput string) (domain.User, error) {
	idInt, err := strconv.Atoi(idInput)
	if err != nil {
		return domain.User{}, err
	}
	value, error := u.Repo.FindUserById(idInt)
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
	idInt, err := strconv.Atoi(idInput)
	if err != nil {
		return err
	}
	value, err := u.Repo.UpdateUser(idInt, user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) CreateProfile(id string, p dto.CreateProfiledto) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	user := domain.User{
		Phone:     p.Phone,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
	value, err := u.Repo.UpdateUser(idInt, user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) BecomeSeller(idInput string) error {
	idInt, err := strconv.Atoi(idInput)
	if err != nil {
		return err
	}
	user := domain.User{
		UserType: "seller",
	}
	value, err := u.Repo.UpdateUser(idInt, user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) RevokeSeller(idInput string) error {
	idInt, err := strconv.Atoi(idInput)
	if err != nil {
		return err
	}
	user := domain.User{
		UserType: "buyer",
	}
	value, err := u.Repo.UpdateUser(idInt, user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
