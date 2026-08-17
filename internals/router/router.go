package router

import (
	"office-expense-management-backend/internals/services"

	"github.com/gin-gonic/gin"
)

func Router(expenseList *services.ExpenseList) *gin.Engine {
	r := gin.Default()

	r.GET("/expense", expenseList.GetExpenseList)
	r.POST("/login", services.Login)

	return r

}
