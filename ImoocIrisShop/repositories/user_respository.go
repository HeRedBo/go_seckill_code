package repositories

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/datamodels"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	Conn() error
	Insert(User *datamodels.User) (userId int64, err error)
}


type UserRepository struct {
	mysqlConn *gorm.DB
}



func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlConn: db}
}

func(u *UserRepository) Conn() (err error) {
	if u.mysqlConn == nil {
		gormMysql , errMysql := common.NewGormMysqlConn()
		if errMysql != nil {
			return errMysql
		}
		u.mysqlConn = gormMysql
	}
	return
}


func (u *UserRepository) Insert(User *datamodels.User) (userId int64, err error) {
	//if err := u.Conn(); err != nil {
	//	return
	//}
	u.mysqlConn.Create(User)
	return User.ID,nil
	//return true
}