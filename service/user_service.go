package service

import (
	"balance-api/model"
	"balance-api/repository"
	"balance-api/utils"
	"net/http"
)

type UserServiceI interface {
	CreateUser() error
	Save(user *model.User) *RestError
	GetBalance(id uint64) (float64, *RestError)
	MakeTransaction(userID uint64, state, amount, transactionID string) *RestError
}

type UserService struct {
	userRepo        repository.UserRepositoryI
	transactionRepo repository.TransctionRepositoryI
}

// NewUserService returns an instance of the UserService
func NewUserService(
	_userRepo repository.UserRepositoryI,
	_transactionRepo repository.TransctionRepositoryI,
) UserServiceI {
	return &UserService{
		userRepo:        _userRepo,
		transactionRepo: _transactionRepo,
	}
}

// CreateUser - creates user
func (srv UserService) CreateUser() error {

	_, err := srv.userRepo.Create(model.User{})

	return err
}

// Save - saves incoming User object
func (service *UserService) Save(user *model.User) *RestError {
	if err := service.userRepo.Save(user); err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}
	return nil
}

// GetBalance - returns user's balance by id
func (service *UserService) GetBalance(id uint64) (float64, *RestError) {
	user, err := service.userRepo.FindByID(id)
	if err != nil {
		return 0, &RestError{
			Status: http.StatusNotFound,
			Error:  err,
		}
	}
	return user.Balance, nil
}

// MakeTransaction - saves transaction
func (service *UserService) MakeTransaction(
	userID uint64,
	state,
	amount,
	transactionID string,
) *RestError {
	parsedAmount, err := utils.AtoiFloat64(amount)
	if err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}

	transactionType := model.TransactionTypeAdd
	if state == "lose" {
		transactionType = model.TransactionTypeSub
	}

	tx := service.transactionRepo.NewTx()
	if tx.Error != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}
	defer tx.Rollback()

	if _, err = service.transactionRepo.Create(model.Transaction{
		UserID:     userID,
		Amount:     parsedAmount,
		Type:       transactionType,
		ExternalID: transactionID,
	}, tx); err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}

	chargedAmount := parsedAmount
	if transactionType == model.TransactionTypeSub {
		chargedAmount *= -1
	}

	if err = service.userRepo.UpdateBalance(userID, chargedAmount, tx); err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}

	if err = tx.Commit().Error; err != nil {
		return &RestError{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}

	return nil
}
