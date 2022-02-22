package services

import (
	"ImoocIrisShop/datamodels"
	"ImoocIrisShop/repositories"
)

type IOrderService interface {
	//GetOrderByID(int64) (*datamodels.Order, error)
	GetOrderInfoBy(int64) (map[string]string, error)
	InsertOrder(*datamodels.Order) (int64, error)
	DeleteByID(int64) bool
	UpdateOrder(*datamodels.Order) error
	//GetAllOrder() ([]*datamodels.Order, error)
	// GetAllOrderInfo() (map[int]map[string]string, error)
	GetAllOrderInfo() ([]*datamodels.Order, error)
	InsertOrderByMessage(*datamodels.Message) (int64, error)

	GetOrderDataBySql() (map[int]map[string]string, error)
}

type OrderService struct {
	OrderRepository repositories.IOrderRepository
}

func NewOrderService(repository repositories.IOrderRepository) IOrderService {
	return &OrderService{OrderRepository: repository}
}

func (o *OrderService) GetOrderByID(orderID int64) (*datamodels.Order, error) {
	return o.OrderRepository.SelectByKey(orderID)
}

func (o *OrderService) GetOrderInfoBy(orderID int64) (map[string]string, error) {
	return o.OrderRepository.SelectInfoByKey(orderID)
}

//
func (o *OrderService) InsertOrder(order *datamodels.Order) (orderID int64, err error) {
	return o.OrderRepository.Insert(order)
}

func (o *OrderService) DeleteByID(orderID int64) (isOk bool) {
	return o.OrderRepository.Delete(orderID)
}

func (o *OrderService) UpdateOrder(order *datamodels.Order) error {
	return o.OrderRepository.Update(order)
}

func (o *OrderService) GetAllOrder() ([]*datamodels.Order, error) {
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) GetAllOrderInfo() ([]*datamodels.Order, error) {
	return o.OrderRepository.SelectAllWithInfo()
}

func (o *OrderService) InsertOrderByMessage(message *datamodels.Message) (orderID int64, err error) {
	order := &datamodels.Order{
		UserId:      message.UserID,
		ProductId:   message.ProductID,
		OrderStatus: datamodels.OrderSuccess,
	}
	return o.InsertOrder(order)
}

func (o *OrderService) GetOrderDataBySql() (map[int]map[string]string, error) {
	return o.OrderRepository.GetOrderDataBySql()
}
