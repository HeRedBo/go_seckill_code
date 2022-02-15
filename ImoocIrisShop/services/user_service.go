package services

import (
	"ImoocIrisShop/datamodels"
	"ImoocIrisShop/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	AddUser(user *datamodels.User) (userId int64,err error)
}


type UserService struct {
	UserRepository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository}
}

func (u *UserService) AddUser(user *datamodels.User) (userId int64,err error) {
	pwdByte, errPwd := GeneratePassword(user.HashPassword)
	if errPwd != nil {
		return userId, errPwd
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.Insert(user)
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword),bcrypt.DefaultCost)
}