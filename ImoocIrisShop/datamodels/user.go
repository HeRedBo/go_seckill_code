package datamodels

type User struct {
	Model
	//ID int `gorm:"primary_key" json:"id" form:"ID" sql:"id"`
	NickName string `json:"nickname" form:"nickname" sql:"nickname"`
	UserName string `json:"username" form:"username" sql:"username"`
	HashPassword string `json:"_" form:"password" sql:"password"`
}

// 模型自定义表名
func (User) TableName() string {
	return "users"
}
