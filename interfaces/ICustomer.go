package interfaces

import ("bankinggo/models"
"go.mongodb.org/mongo-driver/bson"
)

type ICustomer interface{
	CreateCustomer(customer *models.Customer)(error)
	DeleteService(filter bson.M)(error)
	UpdateService(filter bson.M,update bson.M)(error)
	FindService(filter bson.M)([] models.Customer,error)
	FindServiceSum(filter bson.M)([] models.Customer,error)
}