package handler

import (
	emulator "github.com/duramash/constanta-emulator-task/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createTransaction(c *gin.Context) {

	var input emulator.TransactionModel
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.services.Transaction.Create(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success!",
	})
}

func (h *Handler) changeStatusOfTransactionById(c *gin.Context) {
	transactionId := c.Param("id")
	transactionIdInt, err := strconv.Atoi(transactionId)

	var input emulator.TransactionModel
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Transaction.ChangeStatus(transactionIdInt, input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "changed successfully",
	})
}

func (h *Handler) getStatusOfTransactionById(c *gin.Context) {
	transactionId := c.Param("id")
	transactionIdInt, err := strconv.Atoi(transactionId)
	if err != nil {
		return
	}

	status, err := h.services.Transaction.GetStatus(transactionIdInt)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}

func (h *Handler) cancelTransactionById(c *gin.Context) {
	transactionId := c.Param("id")
	transactionIdInt, err := strconv.Atoi(transactionId)
	if err != nil {
		return
	}
	err = h.services.Transaction.Cancel(transactionIdInt)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (h *Handler) getTransactionsByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return
	}
	transactions, err := h.services.Transaction.GetByUserId(userIdInt)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"transactions": transactions,
	})
}

func (h *Handler) getTransactionsByUserEmail(c *gin.Context) {
	userEmail := c.Param("user_email")
	transactions, err := h.services.Transaction.GetByUserEmail(userEmail)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"transactions": transactions,
	})
}
