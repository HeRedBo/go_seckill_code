package controllers

import (
	"ImoocIrisShop/services"

	"github.com/kataras/iris/v12"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.ProductSerive
}
