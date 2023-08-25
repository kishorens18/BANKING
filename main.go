package main

import (
	"bankinggo/config"
	"bankinggo/constants"
	"bankinggo/controllers"
	"bankinggo/routes"
	"bankinggo/services"
	"context"
	"fmt"
	"log"

	//	"rest-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initApp(mongoClient *mongo.Client) {
	//Customer Collection
	ctx = context.TODO()
	CustomerCollection := mongoClient.Database(constants.DatabaseName).Collection("customer")
	CustomerService := services.CustomerServiceInit(CustomerCollection, ctx)
	CustomerController := controllers.InitCustomerController(CustomerService)
	routes.CustomerRoute(server, CustomerController)

	//Loan Collection
	ctx = context.TODO()
	LoanCollection := mongoClient.Database(constants.DatabaseName).Collection("loan")
	LoanService := services.LoanServiceInit(LoanCollection, ctx)
	LoanController := controllers.InitLoanController(LoanService)
	routes.Loanroutes(server, LoanController)

	//Transaction Collection
	ctx = context.TODO()
	TransactionCollection := mongoClient.Database(constants.DatabaseName).Collection("transaction")
	TransactionService := services.TransactionServiceInit(mongoClient, CustomerCollection, TransactionCollection, ctx)
	TransactionController := controllers.InitTransactionController(TransactionService)
	routes.Transactionroutes(server, TransactionController)
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}

	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
