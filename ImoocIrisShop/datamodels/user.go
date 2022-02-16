package datamodels

type User struct {
	Model
	//ID           int64  `gorm:"primary_key" json:"id" form:"ID" sql:"id"`
	Nickname string `json:"nickname" form:"nickname" sql:"nickname"`
	Username string `json:"username" form:"username" sql:"username"`
	Password string `json:"_" form:"password" sql:"password"`
}

// 模型自定义表名
func (User) TableName() string {
	return "users"
}
