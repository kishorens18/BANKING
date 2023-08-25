package services

import (
	"context"
	
	"bankinggo/interfaces"
	"bankinggo/models"
	"go.mongodb.org/mongo-driver/mongo"

	
)

type LoanService struct {
	LoanCollection *mongo.Collection
	ctx               context.Context
}

func LoanServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ILoan{
	return &LoanService{ collection , ctx}
}

func (p *LoanService) CreateLoan(user *models.Loan) (error) {
	 
	_, err := p.LoanCollection.InsertOne(p.ctx, &user)

	if err != nil {
		return   err
	}
	return nil

}