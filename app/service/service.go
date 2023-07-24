package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/xvbnm48/go-credit/app/model"
	app "github.com/xvbnm48/go-credit/app/repository"
	"github.com/xvbnm48/go-credit/helper"
)

type (
	walletService interface {
		RegisterCustomer(cmd *model.RegisterCustomer) (string, error)
	}
	walletServiceImpl struct {
		userRepo    app.UserRepo
		walletRepo  app.WalletRepo
		transaction app.TransactionRepo
	}
)

func (s *walletServiceImpl) RegisterCustomer(cmd *model.RegisterCustomer) (helper.ResponseMessage, error) {
	userId := uuid.NewV4().String()

	account := model.UserAccount{
		UserID:      userId,
		Name:        cmd.Name,
		Email:       cmd.Email,
		PhoneNumber: cmd.Phonenumber,
		UserType:    model.CUSTOMER,
	}

	if err := s.userRepo.Save(&account); err != nil {
		response := helper.ResponseMessage{
			Message: "Failed",
			Status:  "Failed",
			Code:    500,
		}
		return response, errors.New("failed to save user account")
	}
	response := helper.ResponseMessage{
		Message: "Success",
		Status:  "Success",
		Code:    200,
	}
	return response, nil
}
