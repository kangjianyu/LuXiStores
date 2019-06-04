package order_dao


type OrderInfo struct {
	Id 		   uint64		`gorm:"column:id" json:"id"`
	TradeId	   string		`gorm:"column:trade_id" json:"trade_id"`
	Status	   uint64		`gorm:"column:status" json:"status"`
}

func (u *OrderInfo) TableName()string{
	return "mmall_order"
}

type TradeInfo struct {
	Id        uint64  `gorm:"column:id" json:"id"`
	Traded    string  `gorm:"column:trade_id" json:"trade_id"`
	UserId    uint64  `gorm:"column:user_id" json:"user_id"`
	ProductId uint64  `gorm:"column:product_id" json:"product_id"`
	Price     float64 `gorm:"column:price" json:"price"`
	PayType   uint64  `gorm:"column:pay_type" json:"pay_type"`


}
func (u *TradeInfo) TableName()string{
	return "mmall_trade"
}