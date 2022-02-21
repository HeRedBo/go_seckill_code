package repositories

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/datamodels"
	"strconv"

	"github.com/jinzhu/gorm"
)

type IOrderRepository interface {
	Conn() error
	Insert(order *datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectInfoByKey(int64) (map[string]string, error)
	SelectAll() ([]*datamodels.Order, error)
	SelectAllWithInfo() ([]*datamodels.Order, error)
}

type OrderRepository struct {
	table     string
	mysqlConn *gorm.DB
}

func NewOrderRepository(table string, sql *gorm.DB) IOrderRepository {
	return &OrderRepository{
		table:     table,
		mysqlConn: sql,
	}
}

// 数据初始化连接
func (o *OrderRepository) Conn() (err error) {
	if o.mysqlConn == nil {
		mysql, err := common.NewGormMysqlConn()
		if err != nil {
			return err
		}
		o.mysqlConn = mysql
	}
	if o.table == "" {
		o.table = `orders`
	}
	return
}

/**
数据插入
*/
func (o *OrderRepository) Insert(order *datamodels.Order) (order_id int64, err error) {
	//1.判断连接是否存在
	if err = o.Conn(); err != nil {
		return
	}

	o.mysqlConn.Create(order)
	return order.ID, nil

	// 准备SQL
	//sql := "INSERT " + o.table + "SET user_id =?,product_id=?,order_status=?"
	//sql := "INSERT INTO " + o.table + " (user_id,product_id,order_status) values (?,?,?)"
	////sql2 := "INSERT INTO " + o.table + " (user_id,product_id,order_status) values (%d,%d,%d)"
	////fmt.Println("sql info", fmt.Sprintf(sql2, order.UserId, order.ProductId, order.OrderStatus))
	//
	//stmt, errSql := o.mysqlConn.Prepare(sql)
	//if errSql != nil {
	//	return 0, errSql
	//}
	////3.传入参数
	//result, errStmt := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	//if errStmt != nil {
	//	return 0, errStmt
	//}
	//return result.LastInsertId()
}

/**
数据删除操作
*/
func (o *OrderRepository) Delete(orderID int64) (isOk bool) {
	if err := o.Conn(); err != nil {
		return
	}
	result := o.mysqlConn.Where("id = ?", strconv.FormatInt(orderID, 10)).Delete(datamodels.Order{})
	if result.Error != nil {
		return false
	}
	return true
	//return result.Error
	//
	//sql := "DELETE FROM " + o.table + " WHERE id =?"
	//stmt, errStmt := o.mysqlConn.Prepare(sql)
	//if errStmt != nil {
	//	return
	//}
	//_, err := stmt.Exec(orderID)
	//if err != nil {
	//	return false
	//}
	//return true
}

func (o *OrderRepository) Update(order *datamodels.Order) error {
	// 判断链接是否存在
	if err := o.Conn(); err != nil {
		return err
	}

	result := o.mysqlConn.Model(&datamodels.Order{}).Where("id = ?", strconv.FormatInt(order.ID, 10)).Update(order)
	return result.Error

	//sql := `UPDATE ` + o.table + ` set user_id=?,product_id=?,order_status=? where id=` + strconv.FormatInt(order.ID, 10)
	//stmt, err := o.mysqlConn.Prepare(sql)
	//if err != nil {
	//	return err
	//}
	//_, err = stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	//if err != nil {
	//	return err
	//}
	//return nil
}

//// 根据商品ID 查询商品
func (o *OrderRepository) SelectByKey(orderID int64) (Order *datamodels.Order, err error) {
	// 判断链接是否存在
	if err := o.Conn(); err != nil {
		return &datamodels.Order{}, err
	}

	//var Order *datamodels.User
	err = o.mysqlConn.Where(" WHERE id = ?", strconv.FormatInt(orderID, 10)).First(Order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return Order, nil

	//sql := "SELECT * FROM " + o.table + " WHERE id = " + strconv.FormatInt(orderID, 10)
	//row, errRow := o.mysqlConn.Query(sql)
	//defer row.Close()
	//
	//if errRow != nil {
	//	return &datamodels.Order{}, errRow
	//}
	//result := common.GetResultRow(row)
	//if len(result) == 0 {
	//	return &datamodels.Order{}, nil
	//}
	//orderResult = &datamodels.Order{}
	//common.DataToStructByTagSql(result, orderResult)
	//return
}

func (o *OrderRepository) SelectInfoByKey(orderID int64) (orderResult map[string]string, err error) {
	if err := o.Conn(); err != nil {
		return map[string]string{}, err
	}

	var order datamodels.Order
	var table = order.TableName()
	//
	//type Result struct {
	//	 ID int
	//	 UserID int
	//	 ProductID int
	//	 OrderStatus int
	//	 ProductName string
	//	 ProductImages string
	//}

	sql := "SELECT orders.id,orders.user_id, orders.product_id, orders.order_status, product.product_name,  product.product_image FROM " + table + " INNER JOIN product ON " + table + ".product_id = product.id WHERE orders.id = " + strconv.FormatInt(orderID, 10)
	rows, errRow := o.mysqlConn.Raw(sql).Rows()
	defer rows.Close()

	//row, errRow := o.mysqlConn.Query(sql)
	//defer row.Close()
	if errRow != nil {
		return map[string]string{}, errRow
	}
	result := common.GetResultRow(rows)
	if len(result) == 0 {
		return map[string]string{}, nil
	}
	return result, err

}

//
func (o *OrderRepository) SelectAll() (orderArray []*datamodels.Order, errProduct error) {
	//1.判断连接是否存在
	if err := o.Conn(); err != nil {
		return nil, err
	}

	var results []datamodels.Order
	o.mysqlConn.Find(&results)

	for i := 0; i < len(results); i++ {
		tmp := &results[i]
		orderArray = append(orderArray, tmp)
	}

	//sql := "Select * from " + o.table
	//rows, err := o.mysqlConn.Query(sql)
	//
	//defer rows.Close()
	//if err != nil {
	//	return nil, err
	//}
	//
	//result := common.GetResultRows(rows)
	//if len(result) == 0 {
	//	return nil, nil
	//}
	//
	//for _, v := range result {
	//	order := &datamodels.Order{}
	//	common.DataToStructByTagSql(v, order)
	//	orderArray = append(orderArray, order)
	//}

	return orderArray, nil
}

func (o *OrderRepository) SelectAllWithInfo() (orders []*datamodels.Order, err error) {

	if err := o.Conn(); err != nil {
		return nil, err
	}
	//var orders []*datamodels.Order
	err = o.mysqlConn.Preload("Product").Preload("User").Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return orders, nil

	//sql := `select orders.id,product.product_name, product.product_image, orders.order_status from orders left join product on orders.product_id = product.id`
	//rows, errRows := o.mysqlConn.Query(sql)
	//if errRows != nil {
	//	return nil, errRows
	//}
	//return common.GetResultRows(rows), err
}
