package model

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	DefaultAvatar = "https://tse1-mm.cn.bing.net/th/id/R-C.a2b5261ff70a93dfef6221558060fb17?rik=N0U3%2fidX%2fPomwg&riu=http%3a%2f%2fimg.keaitupian.cn%2fuploads%2fupimg%2f1597372353904483.jpg&ehk=pYMmjP6NdHwu99fVm3%2bFPiFcOzTzzt2GT48g5%2bEPaos%3d&risl=&pid=ImgRaw&r=0"
)

// Register 用户注册功能
func (param *UserReq) Register() (*User, string) {
	var user User
	count := 0
	DB.Model(&User{}).Where("username=?", param.Username).First(&user).Count(&count)
	if count != 0 {
		return nil, "用户名已存在"
	}
	user.Username = param.Username
	user.Password = param.Password
	user.Avatar = DefaultAvatar
	DB.Create(&user)
	return &user, "成功"

}

func (param *UserReq) Login() (*User, string) {
	var user User
	err := DB.Where("username=?", param.Username).Find(&user).Error

	if err == gorm.ErrRecordNotFound {
		return nil, "该用户不存在"
	}
	if err != nil {
		return nil, "查询出错了"
	} else {
		if param.Password != user.Password {
			return nil, "密码错误"
		} else {
			return &user, "成功"
		}
	}
}

func SelectUserById(id int) (*User, string) {
	var user User
	err := DB.Where("id=?", id).Find(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "用户不存在"
		} else {
			return nil, "查询出错了"
		}
	} else {

		return &user, "成功"
	}

}

var _users = make(map[int]*User)
var lock = &sync.Mutex{}

func GetUserById(id int) (*User, error) {

	lock.Lock()
	user, ok := _users[id]
	lock.Unlock()
	if ok {
		return user, nil
	}

	var user2 = User{}
	err := DB.Where("id=?", id).Find(&user2).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		} else {
			return nil, err
		}
	}
	lock.Lock()
	_users[id] = &user2
	lock.Unlock()
	return &user2, nil

}
