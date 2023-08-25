package models

import (
	"time"
)

type Transaction struct {
	Transaction_ID int       `json:"transaction_id" bson:"transaction_id"`
	From_ID        int       `json:"from_id" bson:"from_id"`
	To_ID          int       `json:"to_id" bson:"to_id"`
	Amount         int       `json:"amount" bson:"amount"`
	Timestamp      time.Time `json:"timestamp" bson:"timestamp"`
}
