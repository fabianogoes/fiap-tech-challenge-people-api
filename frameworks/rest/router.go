package rest

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	customerHandler *CustomerHandler,
	attendantHandler *AttendantHandler,
) (*Router, error) {
	router := gin.Default()

	router.GET("/", Welcome)
	router.GET("/health", Health)
	router.GET("/env", Environment)

	customers := router.Group("/customers")
	{
		customers.GET("/", customerHandler.GetCustomers)
		customers.GET("/:id", customerHandler.GetCustomer)
		customers.GET("/cpf/:cpf", customerHandler.GetCustomerByCPF)
		customers.POST("/", customerHandler.CreateCustomer)
		customers.PUT("/:id", customerHandler.UpdateCustomer)
		customers.DELETE("/:id", customerHandler.DeleteCustomer)
		customers.POST("/sign-in", customerHandler.SignIn)
	}

	attendants := router.Group("/attendants")
	{
		attendants.GET("/", attendantHandler.GetAttendants)
		attendants.GET("/:id", attendantHandler.GetAttendant)
		attendants.POST("/", attendantHandler.CreateAttendant)
		attendants.PUT("/:id", attendantHandler.UpdateAttendant)
		attendants.DELETE("/:id", attendantHandler.DeleteAttendant)
	}

	return &Router{
		router,
	}, nil
}
