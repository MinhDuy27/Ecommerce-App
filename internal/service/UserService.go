package service

import (
	"errors"

	"github.com/MinhDuy27/Ecommerce-App/domain"
	"github.com/MinhDuy27/Ecommerce-App/internal/dto"
	"github.com/MinhDuy27/Ecommerce-App/internal/helper"
	"log"
	"time"
	"github.com/MinhDuy27/Ecommerce-App/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (u *UserService) SignUp(input dto.SignUpdto) (string, error) {
	hashedPassword, err := u.Auth.HashPassword(input.Password)
	if err != nil {
		return "", err
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
		return "", err
	}
	return token, nil
}

func (u *UserService) Login(input dto.Logindto) (string, error) {
	value, err := u.Repo.FindUserByEmail(input.Email)
	if err != nil {
		return "", err
	}
	err = u.Auth.VerifyPassword(value.Password, input.Password)
	if err != nil {
		return "", err
	}
	token, err := u.Auth.GenerateToken(value.ID, value.Email, value.UserType)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (u *UserService) GetProfilesByID(id uint) (domain.User, error) {
	value, error := u.Repo.FindUserById(id)
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
func (u *UserService) UpdateUser(id uint, user domain.User) error {
	value, err := u.Repo.UpdateUser(uint(id), user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) CreateProfile(id uint, p dto.CreateProfiledto) error {
	user := domain.User{
		Phone:     p.Phone,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
	value, err := u.Repo.UpdateUser(id, user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) BecomeSeller(id uint) error {
	user := domain.User{
		UserType: "seller",
	}
	value, err := u.Repo.UpdateUser(id, user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (u *UserService) RevokeSeller(id uint) error {
	user := domain.User{
		UserType: "buyer",
	}
	value, err := u.Repo.UpdateUser(id, user)
	if err != nil {
		return err
	}
	log.Println(value)
	return nil
}
func (s UserService) IsValidToGetCode(id uint) (int, bool) {

	currentUser, err := s.Repo.FindUserById(id)
	if err != nil {
		log.Fatalln(err)
		return 0, false
	}
	if time.Now().Before(currentUser.Expiry) && currentUser.Verifired {
		return 1, false
	}
	return 0, true
}

func (u *UserService) GetVerificationCode(user domain.User) (uint, error) {
	num, isvalid := u.IsValidToGetCode(user.ID)
	if !isvalid {
		switch num {
		case 0:
			return 0, errors.New("user not found")
		case 1:
			return 0, errors.New("user already verified")
		}
	}
	code, err := helper.GenCode(6)
	if err != nil {
		return 0, err
	}
	update := domain.User{
		Code:   code,
		Expiry: time.Now().Add(30 * time.Minute),
	}
	_, err = u.Repo.UpdateUser(user.ID, update)
	if err != nil {
		return 0, err
	}
	return code, nil
}
func (u *UserService) VerifyCode(id uint, code uint) error {

	num, isvalid := u.IsValidToGetCode(id)
	if !isvalid {
		switch num {
		case 0:
			return errors.New("user not found")
		case 1:
			return errors.New("user already verified")
		}
	}
	user, err := u.Repo.FindUserById(id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}
	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	updateUser := domain.User{
		Verifired: true,
	}

	_, err = u.Repo.UpdateUser(id, updateUser)
	if err != nil {
		return errors.New("unable to verify user")
	}
	return nil
}
