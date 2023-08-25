package routes

import (
	"bankinggo/controllers"

	"github.com/gin-gonic/gin"
	
)

func Loanroutes(router *gin.Engine, controller controllers.LoanController) {
	router.POST("/api/loan/create", controller.CreateLoan)

}
