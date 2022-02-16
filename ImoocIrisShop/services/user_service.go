package services

import (
	"ImoocIrisShop/datamodels"
	"ImoocIrisShop/repositories"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	AddUser(user *datamodels.User) (userId int64, err error)
	IsPwdSuccess(userName string, pwd string) (user datamodels.User, isOK bool)
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository}
}

func (u *UserService) AddUser(user *datamodels.User) (userId int64, err error) {
	pwdByte, errPwd := GeneratePassword(user.Password)
	if errPwd != nil {
		return userId, errPwd
	}
	user.Password = string(pwdByte)
	return u.UserRepository.Insert(user)
}

func (u *UserService) IsPwdSuccess(userName string, pwd string) (user datamodels.User, isOK bool) {
	user2, err := u.UserRepository.FindByName(userName)
	if err != nil {
		return
	}
	fmt.Println(user2)
	isOk, _ := ValidatePassword(pwd, user2.Password)
	if !isOk {
		return datamodels.User{}, false
	}
	return user2, true
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}
