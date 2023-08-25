package interfaces

import ("bankinggo/models")

type ILoan interface{
	CreateLoan(user *models.Loan) (error)
}