package receiver_dao

type GoodsReceiverAddress struct {
	Id 		   uint64		`gorm:"column:id" json:"id"`
	Uid        uint64 		`gorm:"column:uid" json:"uid"`
	Nick       string 		`gorm:"column:nick" json:"nick"`
	Tel        string 		`gorm:"column:tel" json:"tel"`
	Mobile     string 		`gorm:"column:mobile" json:"mobile"`
	Province   string 		`gorm:"column:province" json:"province"`
	City       string 		`gorm:"column:city" json:"city"`
	District   string 		`gorm:"column:district" json:"district"`
	Address    string 		`gorm:"column:address" json:"address"`
	IsDefault  uint8   		`gorm:"column:is_default" json:"is_default"`
}

func (u *GoodsReceiverAddress) TableName()string{
	return "mmall_receiver"
}