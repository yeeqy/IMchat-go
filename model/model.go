package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint `gorm:"primary_key" json:"id" `
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Username string `gorm:"unique" json:"username" `
	Password string
	Avatar   string `gorm:"type:varchar(1024)" json:"avatar" `
}

type Info struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`

	gorm.Model
	FromUserId int   `json:"fromUserId" form:"fromUserId"`
	FromUser   *User `json:"fromUser" form:"fromUser"`

	ToUserId int   `json:"toUserId" form:"toUserId"`
	ToUser   *User `json:"toUser" form:"toUser"`

	Message string `gorm:"type:varchar(1024)" json:"message"`
}

type UserReq struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
