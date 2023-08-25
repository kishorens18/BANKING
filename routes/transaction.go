package routes

import (
	"bankinggo/controllers"

	"github.com/gin-gonic/gin"
)

func Transactionroutes(router *gin.Engine, controller controllers.TransactionController) {
	router.POST("/api/transaction/create", controller.Transfer)

}
