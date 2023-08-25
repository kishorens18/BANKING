package routes

import (
	"bankinggo/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerRoute(router *gin.Engine, controller controllers.CustomerController) {
	router.POST("/api/customer/create", controller.CreateCustomer)
	router.DELETE("/api/customer/delete/:id", controller.DeleteCustomer)
	router.PUT("/api/customer/update/:id", controller.UpdateCustomer)
	router.POST("/api/customer/find", controller.FindCustomer)
}
