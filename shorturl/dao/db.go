package shorturl_dao

import "LuXiStores/common"

var DB iDB = dbimpl{}


type iDB interface {
	AddShortUrl(shortUrl,longUrl string)error
	GetShorUrl(shortUrl string) (ShortUrlInfo,error)
}
type dbimpl struct {

}

func (dbimpl) GetShorUrl(shortUrl string) (ShortUrlInfo, error) {
	tablename := (&ShortUrlInfo{}).TableName()
	info := ShortUrlInfo{}
	err := common.MysqlClient.GetLongUrl(tablename,shortUrl,&info).Error
	return info,err
}

	func (dbimpl) AddShortUrl(shortUrl,longUrl string)error {
	tablename := (&ShortUrlInfo{}).TableName()

	ret := common.MysqlClient.InsertShortUrl(tablename,shortUrl,longUrl)
	return ret.Error
}


