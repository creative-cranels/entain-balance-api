package service

import (
	"balance-api/model"
	"balance-api/repository"
	"net/http"
)

type UserServiceI interface {
	CreateUser() error
	FindByID(id uint64) (*model.User, *RestError)
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

// FindByID - finds user by id and returns it
func (srv UserService) FindByID(id uint64) (*model.User, *RestError) {
	user, err := srv.userRepo.FindByID(id)

	if err != nil {
		return user, &RestError{
			Status: http.StatusNotFound,
			Error:  err,
		}
	}

	return user, nil
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
	return 0, nil
}

// MakeTransaction - saves transaction
func (service *UserService) MakeTransaction(
	state,
	amount,
	transactionID string,
) *RestError {
	return nil
}
