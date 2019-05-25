package user_handler

import (
	"LuXiStores/common"
	"LuXiStores/user/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)
type ProfileData struct {
	Uid uint64
	Nick string
	Token string
	Gender uint8
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
	fmt.Println(userprofile.BirthDate,"出生时间")
	userprofile.BirthDate = uint64(time.Now().Year()) - userprofile.BirthDate
	common.BuildResp(c,userprofile,nil)
	return
}

func UpdateProfile(c *gin.Context){
	data,err:=ioutil.ReadAll(c.Request.Body)
	profiledata := ProfileData{}
	err = json.Unmarshal(data,&profiledata)
	if err!=nil||profiledata.Gender>3{
		fmt.Println(err)
		common.BuildResp(c,nil,common.ErrParam)
		return
	}
	err = user_dao.DB.UpdateUserProfile(profiledata.Uid,profiledata.Nick,profiledata.Gender)
	if err!=nil{
		common.BuildResp(c,nil,common.ErrInternal)
		return
	}
	common.BuildResp(c,nil,nil)
	return
}