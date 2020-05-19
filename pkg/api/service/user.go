package service

import (
	"cmsgo/pkg/api/dao"
	"cmsgo/pkg/api/dto"
	"cmsgo/pkg/api/model"
)

const pwHashBytes = 64

var (
	/* Dao layer */
	userDao = dao.User{}
)

type UserService struct {
}

// Create - create a new account
func (u UserService) Create(userDto dto.UserCreateDto) (*model.User, error) {
	userModel := userDao.GetByUserName(userDto.Username)
	return &userModel, nil
}
