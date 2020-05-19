package dto

import "time"

//UserCreateDto - binding user creation params
type UserCreateDto struct {
	Id          int       `json:"id"`
	Username    string    `form:"username" json:"username" binding:"required"`
	Mobile      string    `form:"mobile" json:"mobile" binding:"required"`
	Sex         int       `form:"sex" json:"sex"`
	Realname    string    `form:"realname" json:"realname"`
	Password    string    `form:"password" json:"password" binding:"required,pwdValidate"`
	Salt        string    `json:"-"`
	Faceicon    string    `json:"faceicon"`
	Email       string    `form:"email" json:"email"`
	Title       string    `form:"title" json:"title"`
	Status      int       `form:"status,default=1" json:"status"`
	CreatedAt   time.Time `type(datetime)" json:"create_at"`
	LastLoginAt time.Time `type(datetime)" json:"-"`
	Roles       string    `form:"roles" json:"roles"`
}
