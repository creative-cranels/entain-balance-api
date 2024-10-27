package handler

import (
	"balance-api/request"
	"balance-api/response"
	"balance-api/service"
	"balance-api/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserServiceI
}

func NewUserHandler(userService service.UserServiceI) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// Balance godoc
// @Summary Get user balance by user_id
// @Description Returns user's balance by user_id
// @ID user-balance
// @Tags User Actions
// @Accept json
// @Param			id				path		int		false	"User ID"
// @Produce json
// @Success 200 {object} response.BalanceResponse
// @Failure 422 {object} response.Error
// @Failure 500 {object} string "Error message"
// @Router /user/{id}/balance [get]
func (h *UserHandler) GetUserBalance(context *gin.Context) {
	rw := utils.GetRequestWrapper(context)
	rw.ParseDefaultPathParams()

	balance, restError := h.UserService.GetBalance(rw.ID)
	if restError != nil {
		response.ErrorResponse(context, restError.Status, restError.Error.Error())
		return
	}

	response.SuccessResponse(context, response.BalanceResponse{
		UserID:  rw.ID,
		Balance: fmt.Sprintf("%.2f", balance),
	})
}

// Balance godoc
// @Summary Get user balance by user_id
// @Description Returns user's balance by user_id
// @ID user-balance
// @Tags User Actions
// @Accept json
// @Param			request			body		request.TransactionRequest	true	"Transaction data"
// @Produce json
// @Success 200 {object} string "message"
// @Failure 422 {object} response.Error
// @Failure 500 {object} string "Error message"
// @Router /user/{id}/transaction [post]
func (h *UserHandler) MakeTransaction(context *gin.Context) {
	rw := utils.GetRequestWrapper(context)
	rw.ParseDefaultPathParams()

	var transactionRequest request.TransactionRequest

	if err := context.ShouldBind(&transactionRequest); err != nil {
		response.ErrorResponse(
			context,
			http.StatusUnprocessableEntity,
			"Invalid request body",
		)
		return
	}

	if err := transactionRequest.Validate(); err != nil {
		response.ErrorResponse(
			context,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	if restError := h.UserService.MakeTransaction(
		rw.ID,
		transactionRequest.State,
		transactionRequest.Amount,
		transactionRequest.TransactionID,
	); restError != nil {
		response.ErrorResponse(context, restError.Status, restError.Error.Error())
		return
	}

	response.SuccessResponse(context, "success")
}
