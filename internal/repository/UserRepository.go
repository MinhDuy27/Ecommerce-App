package repository

import (
	"errors"
	"github.com/MinhDuy27/Ecommerce-App/domain"
	"log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User,error)
	FindUserByEmail(email string) (domain.User,error)
	FindUserById(id uint) (domain.User,error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}
// allow outside get service by function, not directly
func RepositoryImage(db *gorm.DB) userRepository{
	return userRepository{
		Db: db,
	}
}
type userRepository struct{
	Db *gorm.DB
}

func (rp userRepository) CreateUser (u domain.User) (domain.User,error){
	err := rp.Db.Create(&u).Error
	if err != nil{
		log.Printf("create user error:%v",err)
		return domain.User{},errors.New("failed to create user")
	}
	return u,nil
}

func (rp userRepository) FindUserByEmail (email string) (domain.User,error){
	var user domain.User
	err := rp.Db.First(&user,"Email=?",email).Error
	if err != nil{
		log.Printf("find error by email")
		return domain.User{},errors.New("cannot find user")
	}
	return user,nil
}
func (rp userRepository) FindUserById (id uint) (domain.User,error){
	var user domain.User
	err := rp.Db.First(&user,id).Error
	if err != nil{
		log.Printf("find error by ID")
		return domain.User{},errors.New("cannot find user")
	}
	return user,nil
}
func (rp userRepository) UpdateUser (id uint, u domain.User) (domain.User,error){
	var user domain.User
	update_user := map [string]interface{}{
		"FirstName": u.FirstName,
		"LastName": u.LastName,
		"Phone": u.Phone,
	}
	err := rp.Db.Model(&user).Clauses(clause.Locking{Strength: "UPDATE"}).Where("ID=?",id).Updates(update_user).Error
	if err != nil{
		log.Printf("update error")
		return domain.User{},errors.New("cannot update user")
	}
	return user,nil
}