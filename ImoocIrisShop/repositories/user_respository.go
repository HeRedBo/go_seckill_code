package repositories

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/datamodels"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	Conn() error
	Insert(User *datamodels.User) (userId int64, err error)
	FindByName(userName string) (user datamodels.User, err error)
}

type UserRepository struct {
	mysqlConn *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlConn: db}
}

func (u *UserRepository) Conn() (err error) {
	if u.mysqlConn == nil {
		gormMysql, errMysql := common.NewGormMysqlConn()
		if errMysql != nil {
			return errMysql
		}
		u.mysqlConn = gormMysql
	}
	return
}

func (u *UserRepository) Insert(user *datamodels.User) (userId int64, err error) {
	if err2 := u.Conn(); err2 != nil {
		return
	}

	u.mysqlConn.Create(user)
	return user.ID, nil
	//return true
}

func (u *UserRepository) FindByName(userName string) (user datamodels.User, err error) {
	var User datamodels.User
	u.mysqlConn.Select("*").Where("username =?", userName).First(&User)
	if User.ID > 0 {
		return User, nil
	}
	return User, nil
}
