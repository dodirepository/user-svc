package repository

import (
	"github.com/dodirepository/user-svc/infrastructure/database"
	domain "github.com/dodirepository/user-svc/internal/domain/repository"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// AccountRepositoryHandler :nodoc:
func UserRepositoryHandler() domain.UserCrud {
	return &UserRepository{
		db: database.GetConnection(),
	}
}

func (r *UserRepository) GetByID(ID int64) (*domain.Users, error) {
	userData := &domain.Users{}
	query := r.db.Where("ID = ?", ID).First(userData)
	if query.RecordNotFound() {
		return nil, nil
	}
	return userData, query.Error

}
func (r *UserRepository) GetByPhoneOrEmail(req string) (*domain.Users, error) {
	userData := &domain.Users{}
	query := r.db.Where("phone = ? OR email = ?", req, req).First(userData)
	if query.RecordNotFound() {
		return nil, nil
	}
	return userData, query.Error

}
func (r *UserRepository) Create(req domain.Users) (err error) {
	if err := r.db.Create(&domain.Users{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}).Error; err != nil {
		return err
	}

	return nil

}
