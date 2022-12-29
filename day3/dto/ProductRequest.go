package dto

type ProductRequest struct {
	ProductId int32
	Name      string
	Cost      int64
	//composition
	DOP string
}
