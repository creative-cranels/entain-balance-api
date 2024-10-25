package handler

import (
	"balance-api/response"

	"github.com/gin-gonic/gin"
)

type WelcomeHandler struct{}

func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}

// Ping godoc
// @Summary Ping
// @Description Pings the current API
// @ID ping
// @Tags Welcome Actions
// @Accept json
// @Produce json
// @Success 200 {string} string "Pong"
// @Failure 422 {object} response.Error
// @Failure 500 {object} string "Error message"
// @Router /ping [get]
func (h *WelcomeHandler) Ping(context *gin.Context) {
	response.SuccessResponse(context, "Pong")
}
