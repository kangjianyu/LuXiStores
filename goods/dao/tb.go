package goods_dao

type GoodsCategory struct {
	Id uint64 				`gorm:"column:id" json:"id"`
	Name string				`gorm:"column:name" json:"name"`
	ParentId int64			`gorm:"column:parent_id" json:"parent_id"`
	Status int64			`gorm:"column:status" json:"status"`
	SortOrder int64			`gorm:"column:sort_order" json:"sort_order"`
	Key string				`gorm:"column:key" json:"key"`
	Level int64				`gorm:"column:level" json:"level"`
}


type GoodsInfo struct {
	Id uint64				`gorm:"column:id" json:"id"`
	CategoryId int64		`gorm:"column:category_id" json:"category_id"`
	Name string				`gorm:"column:name" json:"name"`
	Subtitle string			`gorm:"column:subtitle" json:"subtitle"`
	MainImage string		`gorm:"column:main_image" json:"main_image"`
	SubImage string			`gorm:"column:sub_image" json:"sub_image"`
	Detail string			`gorm:"column:detail" json:"detail"`
	Price string			`gorm:"column:price" json:"price"`
	Stock string			`gorm:"column:stock" json:"stock"`
	Status string			`gorm:"column:status" json:"status"`
}
type GoodsInfoSlice struct {
	Id uint64				`gorm:"column:id" json:"id"`
	Name string				`gorm:"column:name" json:"name"`
	MainImage string		`gorm:"column:main_image" json:"main_image"`
	Price string			`gorm:"column:price" json:"price"`
	Stock string			`gorm:"column:stock" json:"stock"`
	Status string			`gorm:"column:status" json:"status"`
}
func (u *GoodsCategory) TableName() string {
	return "category"
}

func (u *GoodsInfo) TableName()string{
	return "mmall_product"
}
