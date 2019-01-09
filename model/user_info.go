package model

import (
	"errors"
	"config/server_model"
	"config/util"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type UserInfo struct {
	gorm.Model
	UserName string `json:"user_name"`
	UserId string `json:"user_id"`
}

// 新增用户
func (u *UserInfo)CreateUser(req *server_model.AddUpdateUser) error {
	u.UserName = req.Name
	u.UserId = util.GetMsgId()
	if err := Db.Model(&UserInfo{}).Create(u).Error; err != nil {
		log.Error("create user error", "err", err.Error())
		return errors.New("create user failed")
	}
	return nil
}