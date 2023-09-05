package main

import (
	"context"
	airIotService "github.com/acework2u/air-iot-app-api-service/services/airiot"
	"log"
	"net/http"

	conf "github.com/acework2u/air-iot-app-api-service/config"
	"github.com/acework2u/air-iot-app-api-service/configs"
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/acework2u/air-iot-app-api-service/routers"
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/services/auth"
	clientCog "github.com/acework2u/air-iot-app-api-service/services/clientcoginto"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//Swagger
	_ "github.com/acework2u/air-iot-app-api-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client

	userCollection      *mongo.Collection
	customerCollection  *mongo.Collection
	addressCollection   *mongo.Collection
	deviceCollection    *mongo.Collection
	productCollection   *mongo.Collection
	airThingsCollection *mongo.Collection

	// UserRouterCtl routers.UserRouteController
	UserRouterCtl  handler.UserHandler
	CustomerRouter routers.CustomerController

	// address
	AddressRouter routers.AddressController

	// Device
	DeviceHandler handler.DevicesHandler
	DeviceRouter  routers.DeviceRouter

	//air-iot
	airIoTHandler handler.AirIotHandler
	AirIoTRouter  routers.AirIoTRouter

	//AirThings
	airThingHandler handler.AirThingHandler
	AirThingRouter  routers.AirThingRouter

	//'Client'
	CustService   clientCog.ClientCognito
	ClientHandler handler.ClientHandler
	ClientRouter  routers.ClientController

	//Auth
	AuthRouter   routers.AuthController
	ThingsRouter routers.ThingController

	//Product
	ProductHandler handler.ProductHandler
	ProductRouter  routers.ProductRouter
)

func init() {

	ctx = context.TODO()

	envConf, _ := conf.LoadCongig("")
	// Env
	cognitoRegion := envConf.CognRegion
	cognitoClientId := envConf.CognClientId
	userPoolId := envConf.CognUserPoolId

	mongoclient = configs.ConnectDB()
	userCollection = configs.GetCollection(mongoclient, "user")

	userRepository := repository.NewUserRepositoryDB(userCollection, ctx)
	customerService := service.NewUserService(&userRepository)
	UserRouterCtl = handler.NewUserHandler(&customerService)

	customerCollection = configs.GetCollection(mongoclient, "customers")
	customerRepository := repository.NewCustomerRepositoryDB(customerCollection, ctx)
	custService := service.NewCustomerService(customerRepository)
	custHandler := handler.NewCustomerHandler(&custService)
	CustomerRouter = routers.NewCustomerRouter(custHandler)
	// Address
	addressCollection = configs.GetCollection(mongoclient, "cus_address")
	addrRepo := repository.NewAddressRepositoryDB(addressCollection, ctx)
	addrService := service.NewAddressService(addrRepo)
	addrHandler := handler.NewAddressHandler(addrService)
	AddressRouter = routers.NewAddressRouter(addrHandler)

	// devices
	deviceCollection = configs.GetCollection(mongoclient, "devices")
	deviceRepo := repository.NewDeviceRepositoryDB(ctx, deviceCollection)
	deviceService := service.NewDeviceService(deviceRepo)
	DeviceHandler = handler.NewDeviceHandler(deviceService)
	DeviceRouter = routers.NewDeviceRouter(DeviceHandler)

	//AirThing
	airConfig := &service.AirThingConfig{Region: cognitoRegion, UserPoolId: userPoolId, CognitoClientId: cognitoClientId}
	airThingsCollection = configs.GetCollection(mongoclient, "air_things")
	airRepo := repository.NewAirRepository(ctx, airThingsCollection)
	airThingService := service.NewAirThingService(airConfig, airRepo)
	airThingHandler = handler.NewAirThingHandler(airThingService)
	AirThingRouter = routers.NewAirThingRouter(airThingHandler)

	//AirIoT
	airCfg := &airIotService.AirIoTConfig{
		Region:          cognitoRegion,
		UserPoolId:      userPoolId,
		CognitoClientId: cognitoClientId,
		AirRepo:         airRepo,
	}

	airIoTServ := airIotService.NewAirIoTService(airCfg)
	airIoTHandler = handler.NewAirIoTHandler(airIoTServ)
	AirIoTRouter = routers.NewAirIoTRouter(airIoTHandler)

	//Client
	CustService = clientCog.NewCognitoService(cognitoRegion, cognitoClientId)
	ClientHandler = handler.NewClientHandler(CustService)
	ClientRouter = routers.NewClientRouter(ClientHandler)

	//customerService := service.NewCustomerService(&customerRepository)

	//Auth

	authService := auth.NewCognitoClient(cognitoRegion, userPoolId, cognitoClientId, customerRepository)
	authHandler := handler.NewAuthHandler(authService)
	AuthRouter = routers.NewAuthRouter(authHandler)

	//Thing
	thingService := service.NewThingClient(cognitoRegion, userPoolId, cognitoClientId)
	thingHandler := handler.NewThingsHandler(thingService)
	ThingsRouter = routers.NewThingsRouter(thingHandler)

	//Product
	productCollection = configs.GetCollection(mongoclient, "product")
	productRepo := repository.NewProductRepositoryDB(ctx, productCollection)
	productService := service.NewProductService(productRepo)
	ProductHandler = handler.NewProductHandler(productService)
	ProductRouter = routers.NewProductRouter(ProductHandler)

	server = gin.Default()
	//server = gin.New()

}

// @title Air IoT API Service 2023
// @version 1.0
// @description Air Smart IoT App API Service
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apiKey	BearerAuth
// @in header
// @name Authorization
func main() {

	config, _ := conf.LoadCongig("")

	// DB Connect
	defer mongoclient.Disconnect(ctx)

	startGinServer(config)
}

func startGinServer(config conf.Config) {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Origin}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}
	server.Use(cors.New(corsConfig))
	server.Use(gin.Recovery())

	//server.Use(
	//	middleware.ErrorHandler(),
	//)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "page not found",
		})
	})

	//Production
	router := server.Group("/api/v1")
	// Add Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "OK"})
	})

	//Production
	UserRouterCtl.UserRoute(router)
	CustomerRouter.CustomerRoute(router)
	ClientRouter.ClientRoute(router)
	AuthRouter.AuthRoute(router)
	ThingsRouter.ThingsRoute(router)
	AddressRouter.AddressRoute(router)
	DeviceRouter.DeviceRoute(router)
	AirThingRouter.AirThingRoute(router)
	ProductRouter.ProductRoute(router)
	AirIoTRouter.AirIoTRoute(router)

	// Pro
	//routerPro := server.Group("/api/v2")
	//UserRouterCtl.UserRoute(routerPro)
	//CustomerRouter.CustomerRoute(routerPro)
	//ClientRouter.ClientRoute(routerPro)

	//Pro
	// UserRouterCtl.UserRoute(routePro, UserService)

	// log.Fatal(server.Run(":" + config.Port))
	log.Fatal(server.Run(":8080"))
}
