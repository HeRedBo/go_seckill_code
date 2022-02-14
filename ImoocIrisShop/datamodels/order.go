package datamodels

type Order struct {
	ID int64 `json:"id" sql:"id" imooc:"id"`
	UserId int64 `json:"user_id" sql:"user_id" imooc:"user_id"`
	ProductId int64 `json:"product_id" sql:"product_id"  imooc:"product_id"`
	OrderStatus int64 `json:"order_status" sql:"order_status" imooc:"order_status"`
}


// 模型状态值常量
const (
	OrderWait = iota
	OrderSuccess // 1
	OrderFailed // 2
)