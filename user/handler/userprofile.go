package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)
type ProfileData struct {
	Uid			uint64 `json:"uid"`
	Nick		string `json:"nick"`
	BirthDate	uint64 `json:"birth_date"`
	Gender		uint8  `json:"gender"`
	Token  		string `json:"token"`
}
func GetProfile(c *gin.Context){
	struid := c.Query("uid")
	uid,err:= strconv.Atoi(struid)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrParam)
		return
	}

	userprofile,err := user_dao.DB.GetUserProfileByUid(uint64(uid))
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	userprofile.BirthDate = uint64(time.Now().Year()) - userprofile.BirthDate
	common.BuildResp(c,userprofile,nil)
	return
}

func UpdateProfile(c *gin.Context){
	data,err:=ioutil.ReadAll(c.Request.Body)
	profiledata := ProfileData{}
	err = json.Unmarshal(data,&profiledata)
	if err!=nil||profiledata.Gender>3{
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
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}