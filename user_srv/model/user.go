package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32     `gorm:"primary_key;autoIncrement"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_at"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}

type User struct {
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	Nickname string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'male表示男，female表示女'"`
	Role     int        `gorm:"column:role;default:1;type:tinyint(3) comment '1表示普通用户，2表示管理员'"`
	BaseModel
}
