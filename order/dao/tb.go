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
	TradId    string  `gorm:"column:trade_id" json:"trade_id"`
	UserId    uint64  `gorm:"column:user_id" json:"user_id"`
	ProductId uint64  `gorm:"column:product_id" json:"product_id"`
	Price     float64 `gorm:"column:price" json:"price"`
	PayType   uint64  `gorm:"column:pay_type" json:"pay_type"`
	ReceiverId int64 `gorm:"column:receiver_id" json:"receiver_id"`
	Amount int64	 `gorm:"column:amount" json:"amount"`

}

type OrderComment struct {
	Id			int64	`gorm:"column:id" json:"id"`
	OrderId		int64	`gorm:"column:order_id" json:"order_id"`
	Uid			int64	`gorm:"column:uid" json:"uid"`
	Context		string	`gorm:"column:context" json:"context"`
	Start 		int64	`gorm:"column:start" json:"start"`
}
func (u *TradeInfo) TableName()string{
	return "mmall_trade"
}
func (u *OrderComment) TableName()string{
	return "mmall_order_comment"
}