package services

import (
	"ImoocIrisShop/datamodels"
	"ImoocIrisShop/repositories"
)

type IProductSerive interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product,error)
	DeleteProductByID(int64) bool
	InsertProduct(product *datamodels.Product) (int64,error)
	UpdateProduct(product *datamodels.Product) error
}

type ProductSerive struct {
	productRepository repositories.IProductRepository
}

// 类初始化函数
func NewProductSerive(repository repositories.IProductRepository) *ProductSerive {
	return &ProductSerive{repository}
}


func(p *ProductSerive) GetProductByID(productID int64) (*datamodels.Product,error) {
	return p.productRepository.SelectByKey(productID)
}

// 获取产品数据
func (p *ProductSerive) GetAllProduct() ([]*datamodels.Product,error) {
	return p.productRepository.SelectAll()
}

// 删除产品数据
func(p *ProductSerive) DeleteProductByID(productID int64) bool {
	return p.productRepository.Delete(productID)
}

// 插入产品数据
func(p *ProductSerive) InsertProduct(product *datamodels.Product) (int64,error) {
	return p.productRepository.Insert(product)
}

/**
更新产品数据
 */
func(p *ProductSerive) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}





