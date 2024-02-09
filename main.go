package main

import (
	"context"
	"database/sql"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	airIotService "github.com/acework2u/air-iot-app-api-service/services/airiot"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
	"log"
	"net/http"

	conf "github.com/acework2u/air-iot-app-api-service/config"
	"github.com/acework2u/air-iot-app-api-service/configs"
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/handler/smartapp"
	"github.com/acework2u/air-iot-app-api-service/repository"
	smartAppRepo "github.com/acework2u/air-iot-app-api-service/repository/smartapp"
	"github.com/acework2u/air-iot-app-api-service/routers"
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/services/auth"
	clientCog "github.com/acework2u/air-iot-app-api-service/services/clientcoginto"
	smartAppService "github.com/acework2u/air-iot-app-api-service/services/smartapp"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//Swagger
	_ "github.com/acework2u/air-iot-app-api-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server         *gin.Engine
	ctx            context.Context
	mongoclient    *mongo.Client
	mongoDB2Client *mongo.Client

	Db  *sql.DB
	Db2 *bun.DB
	Db3 *gorm.DB

	userCollection      *mongo.Collection
	customerCollection  *mongo.Collection
	addressCollection   *mongo.Collection
	deviceCollection    *mongo.Collection
	productCollection   *mongo.Collection
	airThingsCollection *mongo.Collection
	scheduleCollection  *mongo.Collection
	// Ac
	productBomCollection *mongo.Collection
	errorCodeCollection  *mongo.Collection

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

	// Jobs
	JobsHandler handler.JobsHandler
	JobsRouter  routers.JobsController

	//Schedule
	ScheduleHandler handler.ScheduleHandle
	ScheduleRouter  routers.ScheduleRouter

	//SmartApp
	AcErrorCodeHandler smartapp.ErrorCodeHandler
	AcErrorCodeRouter  routers.AcErrorCodeRouter
	AcCompHandler      smartapp.CompressorHandler
	AcCompRouter       routers.AcCompressorRouter
)

func init() {

	ctx = context.TODO()

	envConf, _ := conf.LoadCongig("")
	// Env
	cognitoRegion := envConf.CognRegion
	cognitoClientId := envConf.CognClientId
	userPoolId := envConf.CognUserPoolId

	mongoclient = configs.ConnectDB()
	mongoDB2Client = configs.ConnectDB2()

	userCollection = configs.GetCollection(mongoclient, "user")

	// Address
	addressCollection = configs.GetCollection(mongoclient, "cus_address")
	addrRepo := repository.NewAddressRepositoryDB(addressCollection, ctx)
	addrService := service.NewAddressService(addrRepo)
	addrHandler := handler.NewAddressHandler(addrService)
	AddressRouter = routers.NewAddressRouter(addrHandler)

	// User
	userRepository := repository.NewUserRepositoryDB(userCollection, ctx)
	customerService := service.NewUserService(&userRepository)
	UserRouterCtl = handler.NewUserHandler(&customerService)
	// Customer
	customerCollection = configs.GetCollection(mongoclient, "customers")
	customerRepository := repository.NewCustomerRepositoryDB(customerCollection, ctx)
	custService := service.NewCustomerService(customerRepository, addrRepo)
	custHandler := handler.NewCustomerHandler(&custService)
	CustomerRouter = routers.NewCustomerRouter(custHandler)

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

	//Jobs
	jobsService := service.NewJobsService(airConfig)
	jobsHandler := handler.NewJobsHandler(jobsService, thingService)
	JobsRouter = routers.NewJobsController(jobsHandler)

	//scheduleCollection
	scheduleCollection = configs.GetCollection(mongoclient, "job_schedule")
	scheduleRepo := repository.NewScheduleRepository(ctx, scheduleCollection)
	scheduleService := service.NewScheduleService(scheduleRepo, airConfig)
	ScheduleHandler = handler.NewScheduleHandler(scheduleService)
	ScheduleRouter = routers.NewScheduleRouter(ScheduleHandler)

	/*********************** Smart App Service **********/
	//Db, _ = configs.ConnectToMariaDB()
	//Db2, ok := configs.SmartConnect()
	//Db3, ok := configs.Db3Connect()

	//acErrRepo := smartAppRepo.NewAcErrorCodeRepo(Db3)
	//acErrorService := smartAppService.NewAcErrorService(acErrRepo)
	//AcErrorCodeHandler = smartapp.NewErrorCodeHandler(acErrorService)
	//AcErrorCodeRouter = routers.NewAcErrorCodeRouter(AcErrorCodeHandler)

	errorCodeCollection = configs.GetCollection(mongoDB2Client, "error_code")
	acErrRepo := smartAppRepo.NewErrorCodeRepo(ctx, errorCodeCollection)
	acErrService := smartAppService.NewErrorCodeService(acErrRepo)
	AcErrorCodeHandler = smartapp.NewErrorCodeHandler(acErrService)
	AcErrorCodeRouter = routers.NewAcErrorCodeRouter(AcErrorCodeHandler)

	// AcCheck Compressor
	productBomCollection = configs.GetCollection(mongoclient, "product_bom")
	acCompRepo := smartAppRepo.NewBomRepository(ctx, productBomCollection)
	acCompService := smartAppService.NewBomService(acCompRepo)
	AcCompHandler = smartapp.NewCompressorHandler(acCompService)
	AcCompRouter = routers.NewAcCompressorRouter(AcCompHandler)

	_ = scheduleService

	_ = scheduleRepo
	go scheduleService.CornJob()
	// Server Start
	server = gin.Default()
	//server = gin.New()

}

// @title Air IoT API Service 2023
// @version 1.0
// @description Air Smart IoT App API Service
// @Schemes http
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

	server.Use(
		middleware.ErrorHandler(),
	)

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
	JobsRouter.JobsRoute(router)
	ScheduleRouter.ScheduleRoute(router)

	//E-Service
	router2 := server.Group("/smart/app/v1")
	AcErrorCodeRouter.ErrorCodeRoute(router2)
	AcCompRouter.AcCompressorRoute(router2)

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
