package services

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ExpenseList struct {
}

func NewExpenseList() *ExpenseList {
	return &ExpenseList{}
}

func (e *ExpenseList) GetExpenseList(c *gin.Context) {
	data, err := os.ReadFile("dummy.json")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(200, "application/json", data)
}
