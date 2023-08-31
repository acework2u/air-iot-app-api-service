package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	productHandler handler.ProductHandler
}

func NewProductRouter(productHandler handler.ProductHandler) ProductRouter {
	return ProductRouter{productHandler}
}

func (rc *ProductRouter) ProductRoute(rg *gin.RouterGroup) {
<<<<<<< HEAD
	router := rg.Group("/products")
	router.GET("", rc.productHandler.GetProducts)
=======
	router := rg.Group("/product")
>>>>>>> ad1f98be097d983c078b0925f74ee2be200245ae
	router.GET("/:id", rc.productHandler.GetProduct)
	router.POST("", rc.productHandler.PostProduct)
	router.PUT("/:id/info", rc.productHandler.UpdateProduct)
	router.PUT("/warranty", rc.productHandler.UpdateEWarranty)
	router.DELETE("/:id", rc.productHandler.DelProduct)
}
