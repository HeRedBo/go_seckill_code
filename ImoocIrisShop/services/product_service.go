package services

import (
	"ImoocIrisShop/datamodels"
	"ImoocIrisShop/repositories"
)

type IProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(product *datamodels.Product) (int64, error)
	UpdateProduct(product *datamodels.Product) error
	SubNumberOrder(ProductID int64) error
}

type ProductService struct {
	productRepository repositories.IProductRepository
}

// 类初始化函数
func NewProductService(repository repositories.IProductRepository) IProductService {
	return &ProductService{repository}
}

func (p *ProductService) GetProductByID(productID int64) (*datamodels.Product, error) {
	return p.productRepository.SelectByKey(productID)
}

// 获取产品数据
func (p *ProductService) GetAllProduct() ([]*datamodels.Product, error) {
	return p.productRepository.SelectAll()
}

// 删除产品数据
func (p *ProductService) DeleteProductByID(productID int64) bool {
	return p.productRepository.Delete(productID)
}

// 插入产品数据
func (p *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

/**
更新产品数据
*/
func (p *ProductService) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}

func (p *ProductService) SubNumberOrder(ProductID int64) error {
	return p.productRepository.SubProductNum(ProductID)
}
