package internal

import (
	controllers "punto-pago-exposer/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// get services
func PuntoPagoGetServicesHandler(C *gin.Context) {
	coPuntoPago := controllers.ControllerPuntoPagoGetServices{Context: C}
	coPuntoPago.GetListServices()
}

// get bill
func PuntoPagoGetBillHandler(C *gin.Context) {
	coPuntoPago := controllers.ControllerPuntoPagoGetBill{Context: C}
	coPuntoPago.GetBill()
}

// pay bill
func PuntoPagoPayBillHandler(C *gin.Context) {
	coPuntoPago := controllers.ControllerPuntoPagoPayBill{Context: C}
	coPuntoPago.PayBill()
}

func SetupRouterPuntoPago() *gin.Engine {
	router := gin.Default()

	// CORS Setup for KrakenD transactions
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	config.AllowMethods = []string{"GET", "POST"}
	router.Use(cors.New(config))

	groupContainer := router.Group("/punto-pago")
	{
		groupContainer.GET("/get-list-services", PuntoPagoGetServicesHandler)
		groupContainer.GET("/:serviceProvider/bills", PuntoPagoGetBillHandler)
		groupContainer.POST("/:billReference/payments", PuntoPagoPayBillHandler)
	}

	return router
}
