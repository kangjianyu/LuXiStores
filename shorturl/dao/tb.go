package shorturl_dao

type ShortUrlInfo struct {
	Uid        uint64    `gorm:"column:id" json:"id"`
	ShortUrl   string    `gorm:"column:short_url" json:"short_url"`
	LongUrl   string    `gorm:"column:long_url" json:"long_url"`

}

func (u *ShortUrlInfo) TableName() string {
	return "mmall_shorturl"
}
