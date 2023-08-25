package controllers


import (
	"net/http"
	"strings"
	"bankinggo/interfaces"
	"bankinggo/models"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	LoanService interfaces.ILoan
}

func InitLoanController(LoanService interfaces.ILoan) LoanController {
	return LoanController{LoanService} //DI(dependency injection) pattern
}

func  (pc *LoanController) CreateLoan(ctx *gin.Context) {
	var customer *models.Loan
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := pc.LoanService.CreateLoan(customer)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	
}