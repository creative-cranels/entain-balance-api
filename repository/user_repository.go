package repository

import (
	"balance-api/db"
	"balance-api/model"

	"gorm.io/gorm"
)

type UserRepositoryI interface {
	FindByID(id int) (model.User, error)
	Create(user model.User, tx ...*gorm.DB) (uint64, error)
	Save(user *model.User) error
}

type UserRepository struct {
	storage *db.Database
}

func NewUserRepository(db *db.Database) UserRepositoryI {
	return &UserRepository{storage: db}
}

func (repo *UserRepository) FindByID(id int) (model.User, error) {
	var user model.User
	err := repo.storage.Exec(db.IRead).Where("id = ?", id).First(&user).Error
	return user, err
}

func (repo *UserRepository) Create(user model.User, tx ...*gorm.DB) (uint64, error) {
	err := Exec(repo.storage.Exec(db.IWrite), tx).Create(&user).Error
	return user.ID, err
}

func (repo *UserRepository) Save(user *model.User) error {
	return repo.storage.Exec(db.IWrite).Save(user).Error
}
