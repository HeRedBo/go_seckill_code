package datamodels

type Message struct {
	ProductID int64
	UserID    int64
}

// 创建结构体
func NewMessage(UserID, ProductID int64) *Message {
	return &Message{ProductID: ProductID, UserID: UserID}
}
