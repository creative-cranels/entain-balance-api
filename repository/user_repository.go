package repository

import (
	"balance-api/db"
	"balance-api/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryI interface {
	FindByID(id uint64) (*model.User, error)
	Create(user model.User, tx ...*gorm.DB) (uint64, error)
	Save(user *model.User, tx ...*gorm.DB) error
	UpdateBalance(id uint64, amount float64, tx ...*gorm.DB) error
}

type UserRepository struct {
	storage *db.Database
}

func NewUserRepository(db *db.Database) UserRepositoryI {
	return &UserRepository{storage: db}
}

func (repo *UserRepository) FindByID(id uint64) (*model.User, error) {
	var user model.User
	err := repo.storage.Exec(db.IRead).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (repo *UserRepository) Create(user model.User, tx ...*gorm.DB) (uint64, error) {
	err := Exec(repo.storage.Exec(db.IWrite), tx).Create(&user).Error
	return user.ID, err
}

func (repo *UserRepository) Save(user *model.User, tx ...*gorm.DB) error {
	return Exec(repo.storage.Exec(db.IWrite), tx).Save(user).Error
}

func (repo *UserRepository) UpdateBalance(id uint64, amount float64, tx ...*gorm.DB) error {
	result := Exec(repo.storage.Exec(db.IWrite), tx).
		Model(model.User{}).
		Where("id = ? AND balance + ? >= 0", id, amount).
		UpdateColumn("balance", gorm.Expr("balance + ?", amount))
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("couldn't update user balance, negative value appeared")
	}
	return nil
}
