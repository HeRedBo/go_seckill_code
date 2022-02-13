package repositories

import (
	"ImoocIrisShop/backend/common"
	"ImoocIrisShop/datamodels"
	"database/sql"
	"log"
	"strconv"
)

// 第一步、先开发对应的接口
// 第二步、实现定义的接口
type IProductRepository interface {
	// 连接数据
	Conn() (error)
	Insert(product *datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product,error)
}

type productRepository struct {
	table string
	mysqlConn *sql.DB
}

// 初始化函数
func NewProductRepository(table string , db *sql.DB) IProductRepository{
	return &productRepository{table: table, mysqlConn:db}
}



// 连接数据

func (p *productRepository) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = `product`
	}
	return
}

func (p *productRepository) Insert(product *datamodels.Product) (productId int64, err error) {
	//1.判断连接是否存在
	if err=p.Conn();err != nil{
		return
	}
	// 准备SQL
	sql := "INSERT product SET product_name =?,product_num=?,product_image=?,product_url=?"
	stmt, errSql := p.mysqlConn.Prepare(sql)
	if errSql !=nil {
		return 0,errSql
	}
	//3.传入参数
	result,errStmt:=stmt.Exec(product.ProductName,product.ProductNum,product.ProductImage,product.ProductUrl)
	if errStmt !=nil {
		return 0,errStmt
	}
	return result.LastInsertId()
}

/**
商品数据删除
 */
func (p *productRepository) Delete(productID int64) bool {
	// 判断链接是否存在
	if err := p.Conn();err != nil{
		return false
	}
	sql := "DELETE FROM product where id=?"
	stmt,err := p.mysqlConn.Prepare(sql)
	if err!= nil {
		return false
	}
	_,err = stmt.Exec(strconv.FormatInt(productID,10))
	if err != nil {
		return false
	}
	return true
}

func(p *productRepository) Update(product *datamodels.Product) error{
	// 判断链接是否存在
	if err := p.Conn();err != nil{
		return err
	}
	sql := `UPDATE product set product_name=?,product_num=?,product_image=?,product_url=? where id=` + strconv.FormatInt(product.ID,10)
	stmt,err := p.mysqlConn.Prepare(sql)
	if err !=nil {
		return err
	}
	_,err = stmt.Exec(product.ProductName,product.ProductNum,product.ProductImage,product.ProductUrl)
	if err !=nil {
		return err
	}
	return nil
}

// 根据商品ID 查询商品
func(p *productRepository) SelectByKey(productID int64) (productResult *datamodels.Product,err error) {
	// 判断链接是否存在
	if err := p.Conn();err != nil{
		return &datamodels.Product{}, err
	}

	sql := "SELECT * FROM " +p.table+ " WHERE id = " +strconv.FormatInt(productID,10)
	row,errRow :=p.mysqlConn.Query(sql)
	defer row.Close()

	if errRow !=nil {
		return &datamodels.Product{},errRow
	}
	result := common.GetResultRow(row)
	if len(result)==0{
		return &datamodels.Product{},nil
	}
	productResult = &datamodels.Product{}
	common.DataToStructByTagSql(result,productResult)
	return
}

func(p *productRepository) SelectAll() (productArray []*datamodels.Product,errProduct error) {
	//1.判断连接是否存在
	if err:=p.Conn();err!= nil{
		return nil,err
	}
	sql := "Select * from "+p.table
	rows,err := p.mysqlConn.Query(sql)

	defer  rows.Close()
	if err !=nil {
		return nil ,err
	}


	result := common.GetResultRows(rows)
	if len(result)==0{
		return nil,nil
	}

	for _,v :=range result{
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v,product)
		log.Println("product",product)
		productArray=append(productArray, product)
		log.Println("productArray",productArray)

	}

	return productArray,nil
}
