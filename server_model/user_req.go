package server_model

import "errors"

// 添加编辑user
type AddUpdateUser struct {
	Id uint `json:"id"`
	Name string `json:"name"`
}

func (req *AddUpdateUser) CheckParams() error  {
	if req.Name == "" {
		return errors.New("user name is empty")
	}
	return nil
}