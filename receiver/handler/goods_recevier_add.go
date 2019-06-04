package receiver_handler

import (
	"LuXiStores/common"
	receiver_dao "LuXiStores/receiver/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)
type AddGoodsReceiverAddressData struct {
	Uid        uint64 		`json:"uid"`
	Nick       string 		`json:"nick"`
	Tel        string 		`json:"tel"`
	Mobile     string 		`json:"mobile"`
	Province   string 		`json:"province"`
	City       string 		`json:"city"`
	District   string 		`json:"district"`
	Address    string 		`json:"address"`
	IsDefault  uint8   		`json:"is_default"`
}
func AddGoodsReceiverAddress (c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddGoodsReceiverAddressData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.Uid<=0||Data.IsDefault<0||Data.IsDefault>=2{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	if Data.IsDefault==1{
		err := receiver_dao.DB.ChangeDefaultGoodsReceiverAddress(Data.Uid)
		fmt.Println("有到这里")
		err = receiver_dao.DB.AddGoodsReceiverAddress(Data.Uid,Data.Nick,Data.Tel,Data.Mobile,Data.Province,Data.City,Data.District,Data.Address,Data.IsDefault)
		if err!=nil{
			common.BuildResp(c,nil,common.ErrInternal)
			return
		}
		common.BuildResp(c,nil,nil)
		return
	}
	info,err := receiver_dao.DB.GetDefaultGoodsReceiverAddress(Data.Uid)
	if err!=nil||info.Id==0{
		Data.IsDefault = 1
		return
	}
	err = receiver_dao.DB.AddGoodsReceiverAddress(Data.Uid,Data.Nick,Data.Tel,Data.Mobile,Data.Province,Data.City,Data.District,Data.Address,Data.IsDefault)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return


}
