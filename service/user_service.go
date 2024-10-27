package service

import (
	"balance-api/model"
	"balance-api/repository"
	"net/http"
)

type UserServiceI interface {
	CreateUser() error
	Save(user *model.User) *RestError
	GetBalance(id uint64) (float64, *RestError)
	MakeTransaction(state, amount, transactionID string) *RestError
}

type UserService struct {
	userRepo repository.UserRepositoryI
}

// NewUserService returns an instance of the UserService
func NewUserService(
	_userRepo repository.UserRepositoryI,
) UserServiceI {
	return &UserService{
		userRepo: _userRepo,
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
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}
	return user.Balance, nil
}

// MakeTransaction - saves transaction
func (service *UserService) MakeTransaction(
	state,
	amount,
	transactionID string,
) *RestError {
	return nil
}
