package dao

import (
	"cmsgo/pkg/api/model"
	"cmsgo/pkg/common/db"
)

type User struct {
}

// GetByUserName - get user from name
func (u User) GetByUserName(username string) model.User {
	db := db.GetGormDB()
	m := model.User{}
	db.Where("username = ?", username).First(&m)
	return m
}
