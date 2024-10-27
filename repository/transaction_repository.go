package repository

import (
	"balance-api/db"
	"balance-api/model"
	"database/sql"

	"gorm.io/gorm"
)

type TransctionRepositoryI interface {
	NewTx(opts ...*sql.TxOptions) *gorm.DB
	FindByID(id uint64) (*model.Transaction, error)
	FindByExternalID(extID string) (*model.Transaction, error)
	ExistsWithExternalID(extID string) (bool, error)
	Create(transaction model.Transaction, tx ...*gorm.DB) (uint, error)
	Save(transaction *model.Transaction) error
}

type TransctionRepository struct {
	storage *db.Database
}

func NewTransctionRepository(db *db.Database) TransctionRepositoryI {
	return &TransctionRepository{storage: db}
}

func (repo *TransctionRepository) NewTx(opts ...*sql.TxOptions) *gorm.DB {
	return repo.storage.Exec(db.IWrite).Begin()
}

func (repo *TransctionRepository) FindByID(id uint64) (*model.Transaction, error) {
	var transaction model.Transaction
	err := repo.storage.Exec(db.IRead).Where("id = ?", id).First(&transaction).Error
	return &transaction, err
}

func (repo *TransctionRepository) FindByExternalID(extID string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := repo.storage.Exec(db.IRead).Where("external_id = ?", extID).First(&transaction).Error
	return &transaction, err
}

func (repo *TransctionRepository) ExistsWithExternalID(extID string) (bool, error) {
	var exists bool
	err := repo.storage.Exec(db.IRead).
		Model(model.Transaction{}).
		Select("count(*) > 0").
		Where("external_id = ?", extID).
		Find(&exists).Error
	return exists, err
}

func (repo *TransctionRepository) Create(transaction model.Transaction, tx ...*gorm.DB) (uint, error) {
	err := Exec(repo.storage.Exec(db.IWrite), tx).Create(&transaction).Error
	return transaction.ID, err
}

func (repo *TransctionRepository) Save(transaction *model.Transaction) error {
	return repo.storage.Exec(db.IWrite).Save(transaction).Error
}
