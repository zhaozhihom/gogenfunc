package model

const TableNameUser = "users"

type User struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserName string `gorm:"column:user_name;not null" json:"user_name"`
}

func (*User) TableName() string {
	return TableNameUser
}
