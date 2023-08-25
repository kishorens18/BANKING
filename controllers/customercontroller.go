package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"bankinggo/interfaces"
	"bankinggo/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type CustomerController struct {
	CustomerService interfaces.ICustomer
}

func InitCustomerController(CustomerService interfaces.ICustomer) CustomerController {
	return CustomerController{CustomerService} //DI(dependency injection) pattern
}

func (pc *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer *models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := pc.CustomerService.CreateCustomer(customer)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

}

func (pc *CustomerController) DeleteCustomer(ctx *gin.Context) {

	value := ctx.Param("id")

	fmt.Println("value before coverting", value)
	ids, errs := strconv.Atoi(value)
	if errs != nil {
		fmt.Println("ID Number is ", ids)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	filter := bson.M{"customer_id": ids}
	err := pc.CustomerService.DeleteService(filter)
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

}

func (pc *CustomerController) UpdateCustomer(ctx *gin.Context) {
	value := ctx.Param("id")

	fmt.Println("value before coverting", value)
	ids, errs := strconv.Atoi(value)
	if errs != nil {
		fmt.Println("ID Number is ", ids)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}
	filter := bson.M{"customer_id": ids}
	update := bson.M{"$set": bson.M{"customer_name": "warner"}}
	err := pc.CustomerService.UpdateService(filter, update)
	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

}

func (pc *CustomerController) FindCustomer(ctx *gin.Context) {

	// startDate := time.Date(2021, time.August, 1, 0, 0, 0, 0, time.UTC) // Replace with the start date
	// endDate := time.Date(2024, time.August, 31, 23, 59, 59, 999999999, time.UTC) // Replace with the end date

	// Construct the query filter
	// startDate := ctx.Query("start_date")
	// endDate := ctx.Query("end_date")
	// custid:=ctx.Query("customer_id")
	type requestForm struct {
		Startdate string `json:"startdate" bson:"startdate"`
		Enddate   string `json:"enddate" bson:"enddate"`
		Custid    int    `json:"custid" bson:"custid"`
	}
	var requestData requestForm
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s1 := requestData.Startdate
	e1 := requestData.Enddate

	idd := requestData.Custid

	start, err := time.Parse("2006-01-02", s1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}

	end, err := time.Parse("2006-01-02", e1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	filter := bson.M{
		"customer_id": idd,
		"transaction.transaction_date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	results, err1 := pc.CustomerService.FindService(filter)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	for ind, val := range results {
		ctx.JSON(http.StatusOK, gin.H{"index": ind, "value": val})
		fmt.Println(ind, val)

	}

}
