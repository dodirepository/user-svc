package domain

type Users struct {
	ID       int64  `gorm:"column:id;primary_key" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email;unique" json:"email"`
	Phone    string `gorm:"column:phone;unique" json:"phone"`
	Password string `gorm:"column:password" json:"password"`
}

func (Users) TableName() string {
	return "users"
}

type UserCrud interface {
	Create(req Users) error
	GetByID(ID int64) (*Users, error)
	GetByPhoneOrEmail(username string) (*Users, error)
}
