package handler

import (
	"balance-api/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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
	response.SuccessResponse(context, response.BalanceResponse{
		UserID:  1,
		Balance: "10",
	})
}

// Balance godoc
// @Summary Get user balance by user_id
// @Description Returns user's balance by user_id
// @ID user-balance
// @Tags User Actions
// @Accept json
// @Param			id				path		int		false	"User ID"
// @Produce json
// @Success 200 {object} string "message"
// @Failure 422 {object} response.Error
// @Failure 500 {object} string "Error message"
// @Router /user/{id}/balance [get]
func (h *UserHandler) MakeTransaction(context *gin.Context) {
	response.SuccessResponse(context, "success")
}
