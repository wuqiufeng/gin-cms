package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Id          int       `json:"id" example:"1"`
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"username" example:"wutongci"`
	Mobile      string    `json:"mobile" example:"186000000"`
	Sex         int       `json:"sex" example:"1"`
	Realname    string    `json:"realname" example:"黄梧桐"`
	Password    string    `json:"-"`
	Salt        string    `json:"-"`
	Faceicon    string    `json:"faceicon" example:"http://xxx.com"`
	Email       string    `json:"email" example:"xxxx@hotmail.com"`
	Status      int       `json:"status" example:"1"`
	CreatedAt   time.Time `gorm:"type:time;column:create_at;not null;default:CURRENT_TIMESTAMP" json:"created_at,omitempty" example:"2019-07-10 0:39"`
	LastLoginAt time.Time `gorm:"type:time;column:last_login_at;not null;default:CURRENT_TIMESTAMP" json:"logined_at,omitempty" example:"2019-07-10 0:39"`
}

func (User) TableName() string {
	return "user"
}
