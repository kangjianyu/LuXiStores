package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
	log "github.com/jeanphorn/log4go"

)
type ProfileData struct {
	Uid			uint64 `json:"uid"`
	Nick		string `json:"nick"`
	BirthDate	uint64 `json:"birth_date"`
	Gender		uint8  `json:"gender"`
	Token  		string `json:"token"`
}

type AddProfileData struct {
	Uid			uint64 `json:"uid"`
	Nick		string `json:"nick"`
	BirthDate	uint64 `json:"birth_date"`
	Gender		uint8  `json:"gender"`
}
func GetProfile(c *gin.Context){
	struid := c.Query("uid")
	uid,err:= strconv.Atoi(struid)
	if err!=nil{
		log.Warn("input data error :%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	userprofile,err := user_dao.DB.GetUserProfileByUid(uint64(uid))
	if err!=nil{
		log.Warn("get userprofiel error %v uid:%d",err,uid)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	userprofile.BirthDate = uint64(time.Now().Year()) - userprofile.BirthDate
	log.Info("get profile succeed uid:%d",uid)
	common.BuildResp(c,userprofile,nil)
	return

}

func AddProfile(c *gin.Context){
	indata,err := ioutil.ReadAll(c.Request.Body)
	Data := AddProfileData{}
	err = json.Unmarshal(indata,&Data)
	if err!=nil||Data.Uid<=0{
		log.Warn("input data error :%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	err = user_dao.DB.AddUserProfile(Data.Uid,Data.Nick,Data.BirthDate,Data.Gender)
	if err!=nil{
		log.Warn("add userProfile error:%v,Data:%v",err,Data)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info("add userprofiel succeed data:%v",Data)
	common.BuildResp(c,nil,nil)
	return
}
func UpdateProfile(c *gin.Context){
	data,err:=ioutil.ReadAll(c.Request.Body)
	profiledata := ProfileData{}
	err = json.Unmarshal(data,&profiledata)
	if err!=nil||profiledata.Gender>3{
		log.Warn("input data error%v",err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	//token,err :=c.Cookie("sessionid")
	//value,err := user_dao.Rds.GetUserToken(token)
	//if value==""||value!=profiledata.Uid{
	//	common.BuildResp(c,nil,common.ErrRedisKeyNotExist)
	//	return
	//}

	err = user_dao.DB.UpdateUserProfile(profiledata.Uid,profiledata.Nick,profiledata.BirthDate,profiledata.Gender)
	if err!=nil{
		log.Warn("update userprofile error%v,data:%v",err,data)
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	log.Info("update userpfofile succeed data:%v",data)
	common.BuildResp(c,nil,nil)
	return
}