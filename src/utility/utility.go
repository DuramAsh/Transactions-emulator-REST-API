package utility

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorMessage struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorMessage{message})
}

func IsEmail(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return true
		}
	}
	return false
}

func StatusIsAvailable(status string) bool {
	possibleStatuses := []string{"NEW", "SUCCESS", "FAILURE", "ERROR"}
	for i := 0; i < len(possibleStatuses); i++ {
		if status == possibleStatuses[i] {
			return true
		}
	}
	return false
}

func StatusIsTerminal(status string) bool {
	if status == "SUCCESS" || status == "FAILURE" {
		return true
	}
	return false
}
