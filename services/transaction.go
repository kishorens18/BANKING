package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"bankinggo/interfaces"
	"bankinggo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	client                *mongo.Client
	CustomerCollection    *mongo.Collection
	TransactionCollection *mongo.Collection
	ctx                   context.Context
}

func TransactionServiceInit(client *mongo.Client, customer *mongo.Collection, transaction *mongo.Collection, ctx context.Context) interfaces.ITransaction {
	return &TransactionService{client, customer, transaction, ctx}
}

func (p *TransactionService) Transfer(fromid int, toid int, amount int) (string, error) {

	//Create a session
	session, err := p.client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(context.Background())

	//Start a transaction
	_, err = session.WithTransaction(context.Background(), func(ctx mongo.SessionContext) (interface{}, error) {

		//two update queries(dec, inc)
		//deducting from
		filter1 := bson.M{"customer_id": fromid}
		update1 := bson.M{"$inc": bson.M{"balance": -(amount)}}
		_, err1 := p.CustomerCollection.UpdateOne(context.Background(), filter1, update1)
		if err1 != nil {
			fmt.Println("Transaction Failed")
			return nil, err1
		}

		//incrementing to
		filter2 := bson.M{"customer_id": toid}
		update2 := bson.M{"$inc": bson.M{"balance": amount}}
		_, err2 := p.CustomerCollection.UpdateOne(context.Background(), filter2, update2)
		if err2 != nil {
			fmt.Println("Transaction Failed")
			return nil, err2
		}

		//inserting the transaction
		transactionToInsert := models.Transaction{
			Transaction_ID: 1,
			From_ID:        fromid,
			To_ID:          toid,
			Amount:         amount,
			Timestamp:      time.Time{},
		}
		_, err := p.TransactionCollection.InsertOne(context.Background(), transactionToInsert)
		if err != nil {
			fmt.Println("Transaction not inserted")
			return "Transaction not inserted", err
		}
		return "Transaction inserted", nil
	})
	return "hello", nil
}
