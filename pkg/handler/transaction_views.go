package handler

import (
	emulator "github.com/duramash/constanta-emulator-task/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTransaction(c *gin.Context) {
	userId, ok1 := c.Get("user_id")
	userEmail, ok2 := c.Get("user_email")
	amount, ok3 := c.Get("amount")
	currency, ok4 := c.Get("currency")

	if !(ok1 && ok2 && ok3 && ok4) {
		NewErrorResponse(c, http.StatusInternalServerError, "Some body element was not given")
		return
	}

	var input emulator.TransactionModel
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) changeStatusOfTransactionById(c *gin.Context) {

}

func (h *Handler) getStatusOfTransactionById(c *gin.Context) {

}

func (h *Handler) getTransactionsByUserId(c *gin.Context) {

}

func (h *Handler) getTransactionsByUserEmail(c *gin.Context) {

}

func (h *Handler) cancelTransactionById(c *gin.Context) {

}
