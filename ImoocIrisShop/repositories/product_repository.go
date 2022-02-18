package repositories

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/datamodels"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

// 第一步、先开发对应的接口
// 第二步、实现定义的接口
type IProductRepository interface {
	// 连接数据
	Conn() error
	Insert(product *datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
	SubProductNum(productID int64) error
}

type productRepository struct {
	table     string
	mysqlConn *gorm.DB
}

// 初始化函数
func NewProductRepository(table string, db *gorm.DB) IProductRepository {
	return &productRepository{table: table, mysqlConn: db}
}

// 连接数据

func (p *productRepository) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewGormMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	return
}

func (p *productRepository) Insert(product *datamodels.Product) (productId int64, err error) {
	//1.判断连接是否存在
	if err = p.Conn(); err != nil {
		return
	}

	p.mysqlConn.Create(product)
	return product.ID, nil

	// 准备SQL
	//sql := "INSERT product SET product_name =?,product_num=?,product_image=?,product_url=?"
	//stmt, errSql := p.mysqlConn.Prepare(sql)
	//if errSql != nil {
	//	return 0, errSql
	//}
	////3.传入参数
	//result, errStmt := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	//if errStmt != nil {
	//	return 0, errStmt
	//}
	//return result.LastInsertId()
}

/**
商品数据删除
*/
func (p *productRepository) Delete(productID int64) bool {
	// 判断链接是否存在
	if err := p.Conn(); err != nil {
		return false
	}
	p.mysqlConn.Where("id = ?", strconv.FormatInt(productID, 10)).Delete(datamodels.Product{})
	return true

	//sql := "DELETE FROM product where id=?"
	//stmt, err := p.mysqlConn.Prepare(sql)
	//if err != nil {
	//	return false
	//}
	//_, err = stmt.Exec(strconv.FormatInt(productID, 10))
	//if err != nil {
	//	return false
	//}
	//return true
}

func (p *productRepository) Update(product *datamodels.Product) error {
	// 判断链接是否存在
	if err := p.Conn(); err != nil {
		return err
	}

	p.mysqlConn.Model(&datamodels.Product{}).Where("id = ?", strconv.FormatInt(product.ID, 10)).Update(product)

	//sql := `UPDATE product set product_name=?,product_num=?,product_image=?,product_url=? where id=` + strconv.FormatInt(product.ID, 10)
	//stmt, err := p.mysqlConn.Prepare(sql)
	//if err != nil {
	//	return err
	//}
	//_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	//if err != nil {
	//	return err
	//}
	return nil
}

// 根据商品ID 查询商品
func (p *productRepository) SelectByKey(productID int64) (productResult *datamodels.Product, err error) {
	// 判断链接是否存在
	if err := p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}

	var product datamodels.Product
	err = p.mysqlConn.Where("id = ?", strconv.FormatInt(productID, 10)).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &product, nil

	//
	//sql := "SELECT * FROM " + p.table + " WHERE id = " + strconv.FormatInt(productID, 10)
	//row, errRow := p.mysqlConn.Query(sql)
	//defer row.Close()
	//
	//if errRow != nil {
	//	return &datamodels.Product{}, errRow
	//}
	//result := common.GetResultRow(row)
	//if len(result) == 0 {
	//	return &datamodels.Product{}, nil
	//}
	//productResult = &datamodels.Product{}
	//common.DataToStructByTagSql(result, productResult)
	//return
}

func (p *productRepository) SelectAll() (productArray []*datamodels.Product, errProduct error) {
	//1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return nil, err
	}
	var results []datamodels.Product
	p.mysqlConn.Find(&results)

	if len(results) == 0 {
		return nil, nil
	}
	log.Println(results)
	// 遍历数据 指针
	slices := make([]*datamodels.Product, 0, 100)
	for _, v := range results {
		//v = &v
		slices = append(slices, &v)
		productArray = append(productArray, &v)
	}
	log.Println("slices", slices)
	log.Println("productArray", productArray)

	return slices, nil

	//sql := "Select * from " + p.table
	//rows, err := p.mysqlConn.Query(sql)
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
	//	product := &datamodels.Product{}
	//	common.DataToStructByTagSql(v, product)
	//	log.Println("product", product)
	//	productArray = append(productArray, product)
	//	log.Println("productArray", productArray)
	//
	//}
	//
	//return productArray, nil
}

func (p *productRepository) SubProductNum(ProductID int64) error {

	if err := p.Conn(); err != nil {
		return err
	}
	var product datamodels.Product
	p.mysqlConn.Model(&product).Where("id = ? ", strconv.FormatInt(ProductID, 10)).Update("product_num", gorm.Expr("product_num-1"))
	return nil

	//sql := "update " + p.table + " SET " + " product_num=product_num-1 where id = " + strconv.FormatInt(ProductID, 10) + " AND product_num >0 "
	//stmt, err := p.mysqlConn.Prepare(sql)
	//if err != nil {
	//	return err
	//}
	//_, err = stmt.Exec()
	//return err

}
