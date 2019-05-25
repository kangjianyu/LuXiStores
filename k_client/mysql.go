package k_client

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/jinzhu/gorm"
	"runtime"
	"strings"
)

type MysqlClient struct {
	// 包装一层, 方便打日志和统计数据
	DB *gorm.DB

}
func NewMysqlClient(db *gorm.DB) *MysqlClient {
	return &MysqlClient{DB: db}
}

func MysqlLog(ret *gorm.DB) {
	pc, _, _, _ := runtime.Caller(2)
	detail := runtime.FuncForPC(pc)
	funcPath := strings.Split(detail.Name(), "/")
	callerName := funcPath[len(funcPath)-1]
	if ret.Error == nil {
		log.Info("(MYSQL %s)|%s|exist", callerName,ret.Value )
	} else if ret.RecordNotFound()==true {
		log.Info("(MYSQL %s)|%s|not exist", callerName,ret.Value )
	}else{
		log.Error("(REDIS %s)|%s|err:%s",callerName,ret.Error)
	}
}

func (m *MysqlClient) GetByUid(table string,uid uint64,user interface{}) *gorm.DB {
	ret := m.DB.Table(table).Where("uid=?",uid).First(user)
	MysqlLog(ret)
	return ret
}
func (m *MysqlClient) UpdateByProfile(table string,uid uint64,user interface{}) *gorm.DB{
	ret := m.DB.Table(table).Where("uid=?",uid).Updates(user)
	MysqlLog(ret)
	return ret
}
func (m *MysqlClient) Ping() {
	ret := m.DB.Table("userprofile")
	fmt.Println(ret.Error,"连接")
}
