package main

import (
	"context"
	"log"
	"net/http"

	conf "github.com/acework2u/air-iot-app-api-service/config"
	"github.com/acework2u/air-iot-app-api-service/configs"
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/acework2u/air-iot-app-api-service/routers"
	service "github.com/acework2u/air-iot-app-api-service/services"
	clientCog "github.com/acework2u/air-iot-app-api-service/services/clientcoginto"
	services "github.com/acework2u/air-iot-app-api-service/services/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client

	userCollection     *mongo.Collection
	customerCollection *mongo.Collection

	UserService services.UserService
	// UserRouterCtl routers.UserRouteController
	UserRouterCtl  handler.UserHandler
	CustomerRouter routers.CustomerController

	//Client
	ClientHandler handler.ClientHandler
	ClientRouter  routers.ClientController
)

func init() {

	ctx = context.TODO()

	// connect to mongoDB
	// mongoconn := options.Client().ApplyURI(configs.EnvMongoURI)
	// mongoclient, err := mongo.Connect(ctx, mongoconn)
	// if err != nil {
	// 	panic(err)
	// }
	// if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("MongoDB Successfull connected...")

	mongoclient = configs.ConnectDB()
	userCollection = configs.GetCollection(mongoclient, "user")
	// authCollection = configs.GetCollection(mongoclient, "user")
	// userRepo = repository.NewUserRepositoryDB(authCollection, ctx)
	// UserService = services.NewUserService(userRepo)

	userRepository := repository.NewUserRepositoryDB(userCollection, ctx)
	customerService := service.NewUserService(&userRepository)
	UserRouterCtl = handler.NewUserHandler(&customerService)

	customerCollection = configs.GetCollection(mongoclient, "customers")
	customerRepository := repository.NewCustomerRepositoryDB(customerCollection, ctx)
	custService := service.NewCustomerService(customerRepository)
	custHandler := handler.NewCustomerHandler(&custService)
	CustomerRouter = routers.NewCustomerRouter(custHandler)

	_ = custService

	//Client
	cognitoRegion := "ap-southeast-1"
	cognitoClientId := "qq74q62sm1jfg8t7qetmo3a86"
	clientService := clientCog.NewCognitoService(cognitoRegion, cognitoClientId)
	ClientHandler = handler.NewClientHandler(clientService)
	ClientRouter = routers.NewClientRouter(ClientHandler)

	//customerService := service.NewCustomerService(&customerRepository)

	server = gin.Default()

}

func main() {
	config, err := conf.LoadCongig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	//fmt.Print(config)

	defer mongoclient.Disconnect(ctx)
	startGinServer(config)
}

func startGinServer(config conf.Config) {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Origin}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "page not found",
		})
	})

	router := server.Group("/api")

	// routePro := server.Group("/api/v2")
	// router.GET("/healthchecker", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "OK"})
	// })

	//Uat
	UserRouterCtl.UserRoute(router)
	CustomerRouter.CustRoute(router)
	ClientRouter.ClientRoute(router)

	//Pro
	// UserRouterCtl.UserRoute(routePro, UserService)

	// log.Fatal(server.Run(":" + config.Port))
	log.Fatal(server.Run(":3000"))
}
