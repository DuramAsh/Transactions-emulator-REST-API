package handler

import (
	emulator "github.com/duramash/constanta-emulator-task/src/models"
	"github.com/duramash/constanta-emulator-task/src/utility"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) login(c *gin.Context) {
	var credentials emulator.Credentials

	if err := c.BindJSON(&credentials); err != nil {
		utility.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := utility.GenerateToken(credentials.Username, credentials.Password)
	if err != nil {
		utility.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
