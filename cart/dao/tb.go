package cart_dao


type GoodsCartInfo struct {
	Id	uint64				`gorm:"column:id" json:"id"`
	UserId	uint64			`gorm:"column:user_id" json:"user_id"`
	ProductId	uint64		`gorm:"column:product_id" json:"product_id"`
	Quantity	uint64		`gorm:"column:quantity" json:"quantity"`
}
func (u *GoodsCartInfo) TableName()string{
	return "mmall_cart"
}