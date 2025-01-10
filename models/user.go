package models

type User struct {
	ID    int64  `orm:"pk;auto;column(id)" json:"id"`
	Name  string `orm:"size(128);column(name)" json:"name" valid:"Required"`
	Age   int    `orm:"column(age)" json:"age" valid:"Required"`
	Email string `orm:"size(128);unique;column(email)" json:"email" valid:"Required;Email"`
}

func (u *User) TableName() string {
	return "users"
}
