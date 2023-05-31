package models

import "strconv"

type User struct {
	ID
	Email    string `json:"name" gorm:"size:30;not null;comment:用户邮箱"`
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
