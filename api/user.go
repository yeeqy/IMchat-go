package api

import (
	"IM-chat/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserRegister(c *gin.Context) {
	var userReq *model.UserReq
	if err := c.ShouldBind(&userReq); err == nil {
		resp, msg := userReq.Register()
		if msg == "成功" {
			user := &User{
				Id:        int(resp.ID),
				Username:  resp.Username,
				UpdatedAt: resp.UpdatedAt.String(),
				Avatar:    resp.Avatar,
			}
			data := &UserData{
				User:  user,
				Token: MD5(user.Username),
			}
			c.JSON(200, UserResp{
				Flag:  true,
				Data:  data,
				Msg:   msg,
				Error: "",
			})
		} else {
			c.JSON(200, UserResp{
				Flag:  false,
				Data:  "nil",
				Msg:   msg,
				Error: "",
			})
		}
	} else {
		c.JSON(400, UserResp{
			Flag:  true,
			Data:  "nil",
			Msg:   "参数绑定问题",
			Error: err.Error(),
		})
	}
}
func UserLogin(c *gin.Context) {
	var user *model.UserReq
	if err := c.ShouldBind(&user); err == nil {
		resp, msg := user.Login()
		if msg == "成功" {
			user := &User{
				Id:        int(resp.ID),
				Username:  resp.Username,
				UpdatedAt: resp.UpdatedAt.String(),
				Avatar:    resp.Avatar,
			}
			data := &UserData{
				User:  user,
				Token: MD5(user.Username),
			}
			c.JSON(200, UserResp{
				Flag:  true,
				Data:  data,
				Msg:   msg,
				Error: "",
			})
		} else {
			c.JSON(200, UserResp{
				Flag:  false,
				Data:  "nil",
				Msg:   msg,
				Error: "",
			})
		}
	} else {
		c.JSON(400, UserResp{
			Flag:  true,
			Data:  "nil",
			Msg:   "参数绑定问题",
			Error: err.Error(),
		})
	}

}

func UserSelect(c *gin.Context) {
	userId := c.Query("userId")
	id, _ := strconv.Atoi(userId)
	resp, msg := model.SelectUserById(id)
	if msg == "成功" {
		user := &User{
			Id:        int(resp.ID),
			Username:  resp.Username,
			UpdatedAt: resp.UpdatedAt.String(),
			Avatar:    resp.Avatar,
		}
		data := &UserData{
			User:  user,
			Token: MD5(user.Username),
		}
		c.JSON(200, UserResp{
			Flag:  true,
			Data:  data,
			Msg:   msg,
			Error: "",
		})
	} else {
		c.JSON(200, UserResp{
			Flag:  false,
			Data:  "nil",
			Msg:   msg,
			Error: "",
		})
	}

}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func GetUser(id int) *User {
	var user User
	model.DB.Where("id=?", id).Find(&user)
	return &user
}
