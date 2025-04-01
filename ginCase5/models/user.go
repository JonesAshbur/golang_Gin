package models

type User struct { //default table name is users
	Id       int
	UserName string `gorm:"column:username"`
	Age      int
	Email    string
	AddTime  int
}

func (user User) TableName() string {
	return "user"
}
