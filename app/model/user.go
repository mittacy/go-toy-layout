package model

import "time"

const (
	UserIsDeletedNo  = 0
	UserIsDeletedYes = 1
)

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	IsDeleted int8      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*User) TableName() string {
	return "table_name"
}

