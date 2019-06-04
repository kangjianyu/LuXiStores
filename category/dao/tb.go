package category_dao

type GoodsCategory struct {
	Id uint64 				`gorm:"column:id" json:"id"`
	Name string				`gorm:"column:name" json:"name"`
	ParentId int64			`gorm:"column:parent_id" json:"parent_id"`
	Status int64			`gorm:"column:status" json:"status"`
	SortOrder int64			`gorm:"column:sort_order" json:"sort_order"`
	Key string				`gorm:"column:key" json:"key"`
	Level int64				`gorm:"column:level" json:"level"`
}

func (u *GoodsCategory) TableName() string {
	return "category"
}