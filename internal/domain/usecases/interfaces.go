package domain

type UserUsecaseInterface interface {
	Login(UserLogin) (interface{}, *ErrorResponse)
	Insert(UserCreate) (err *ErrorResponse)
	Detail(ID int64) (interface{}, *ErrorResponse)
}
