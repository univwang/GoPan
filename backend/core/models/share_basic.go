package models

import "time"

type ShareBasic struct {
	Id                 int       `json:"id"`
	Identity           string    `json:"identity"`
	UserIdentity       string    `json:"user_identity"`
	RepositoryIdentity string    `json:"repository"`
	ExpiredTime        int       `json:"expired_time"`
	ClickNum           int       `json:"click_num"`
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
