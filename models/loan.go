package models


type Loan struct{

	Loan_Amount int `json:"loan_amount" bson:"loan_amount"`
	Customer_ID int  `json:"customer_id" bson:"customer_id"`
	
}