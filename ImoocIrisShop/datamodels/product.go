package datamodels

type Product struct {
	ID int64 `json:"id" sql:"id", imooc:"id"`
	ProductName  string `json:"product_name"  sql:"product_name" imooc:"product_name"`
	ProductNum  int64 `json:"product_num" sql:"product_num" imooc:"product_num"`
	ProductImage  string `json:"product_image" sql:"product_image" imooc:"product_image"`
	ProductUrl  string `json:"product_url" sql:"product_url" imooc:"product_url"`
}
