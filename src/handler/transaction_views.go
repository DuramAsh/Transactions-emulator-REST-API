package handler

import (
	emulator "github.com/duramash/constanta-emulator-task/src/models"
	"github.com/duramash/constanta-emulator-task/src/utility"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createTransaction(c *gin.Context) {

	var input emulator.TransactionModel
	err := c.BindJSON(&input)
	if err != nil {
		utility.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Transaction.Create(input)
	if err != nil {
		utility.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "created successfully",
	})
}

func (h *Handler) changeStatusOfTransactionById(c *gin.Context) {

	err := utility.IdentifyUser(c)

	if err != nil {
		utility.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	transactionId := c.Param("id")
	transactionIdInt, err := strconv.Atoi(transactionId)
	if err != nil {
		utility.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var body emulator.Status
	if err := c.BindJSON(&body); err != nil {
		utility.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.services.Transaction.ChangeStatus(transactionIdInt, body)
	if err != nil {
		utility.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"new status": status,
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
		utility.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
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
		utility.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (h *Handler) getTransactionsByUser(c *gin.Context) {
	userAttribute := c.Param("user_attr")

	var userId int
	var transactions []emulator.TransactionModel
	var err error

	if !utility.IsEmail(userAttribute) {
		userId, _ = strconv.Atoi(userAttribute)
		transactions, err = h.services.Transaction.GetByUserId(userId)
	} else {
		transactions, err = h.services.Transaction.GetByUserEmail(userAttribute)
	}

	if err != nil {
		utility.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"transactions": transactions,
	})
}
