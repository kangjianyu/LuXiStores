package goods_dao



type GoodsInfo struct {
	Id uint64				`gorm:"column:id" json:"id"`
	CategoryId int64		`gorm:"column:category_id" json:"category_id"`
	Name string				`gorm:"column:name" json:"name"`
	Subtitle string			`gorm:"column:subtitle" json:"subtitle"`
	MainImage string		`gorm:"column:main_image" json:"main_image"`
	SubImage string			`gorm:"column:sub_image" json:"sub_image"`
	Detail string			`gorm:"column:detail" json:"detail"`
	Price float64			`gorm:"column:price" json:"price"`
	Stock string			`gorm:"column:stock" json:"stock"`
	Status string			`gorm:"column:status" json:"status"`
}

type GoodsInfoSlice struct {
	Id uint64				`gorm:"column:id" json:"id"`
	Name string				`gorm:"column:name" json:"name"`
	MainImage string		`gorm:"column:main_image" json:"main_image"`
	Price float64			`gorm:"column:price" json:"price"`
	Stock string			`gorm:"column:stock" json:"stock"`
	Status string			`gorm:"column:status" json:"status"`
}



func (u *GoodsInfo) TableName()string{
	return "mmall_product"
}
