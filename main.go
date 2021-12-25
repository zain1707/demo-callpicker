package main

import (
	"context"
	"errors"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"

	_inboundHttpDelivery "github.com/gomarkho/demo-callpicker/inbound/delivery/http"
	_inboundRepository "github.com/gomarkho/demo-callpicker/inbound/repository"
	_inboundUseCase "github.com/gomarkho/demo-callpicker/inbound/usecase"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	// check if required env variables are provided
	hasValidEnvVariables()
	// database settings
	dbClient := connectDatabaseClient()
	dbConn := dbClient.Database(os.Getenv("DB_DEFAULT"))
	defer closeDatabaseClient(dbClient)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	pub := e.Group("/v1")
	priv := e.Group("/v1")
	////Init mongo db connection
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	//// services
	//s3 := service.S3Service{}
	//w := service.WebSocketService{}
	//
	//pub.GET("/websocket", w.WebsocketConnection)
	//
	//ideaRepo := _ideaRepo.NewIdeaRepository(dbConn, neo)
	//
	//userRepo := _userRepo.NewUserRepository(dbConn)
	//
	//
	//likeNotificationRepo := _likeNotificationRepo.NewLikeNotificationRepository(dbConn)

	inboundRepo := _inboundRepository.NewLikeRepository(dbConn)
	inboundUseCase := _inboundUseCase.NewInboundUseCase(inboundRepo)
	_inboundHttpDelivery.NewInboundHandler(pub, priv, inboundUseCase)

	if os.Getenv("ENV") == "prod" {
		// Server
		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
		//e.Logger.Fatal(e.StartTLS(":"+os.Getenv("PORT"), os.Getenv("CERT_PATH"), os.Getenv("CERT_KEY_PATH")))

	} else {
		// Server
		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	}

}

func connectDatabaseClient() *mongo.Client {
	ctx := context.Background()

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("DB_HOST")).SetAuth(options.Credential{
		AuthSource: os.Getenv("DB_AUTH_DB"), Username: os.Getenv("DB_USERNAME"), Password: os.Getenv("DB_PASSWORD"),
	}))

	if err != nil {
		logrus.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	return client
}

func closeDatabaseClient(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}
}

func hasValidEnvVariables() {
	if os.Getenv("PORT") == "" {
		panic(errors.New("Please provide valid PORT"))
	}
	if os.Getenv("ENV") == "" {
		panic(errors.New("Please provide valid ENV"))
	}

	//if os.Getenv("CERT_KEY_PATH") == "" {
	//	panic(errors.New("Please provide valid CERT_KEY_PATH"))
	//	}

	//	if os.Getenv("CERT_PATH") == "" {
	//		panic(errors.New("Please provide valid CERT_PATH"))
	//	}
	//if os.Getenv("LOG_LEVEL") == "" {
	//	panic(errors.New("Please provide valid LOG_LEVEL"))
	//}
	if os.Getenv("DB_HOST") == "" {
		panic(errors.New("Please provide valid DB_HOST"))
	}
	if os.Getenv("DB_DEFAULT") == "" {
		panic(errors.New("Please provide valid DB_DEFAULT"))
	}

	if os.Getenv("DB_USERNAME") == "" {
		panic(errors.New("Please provide valid DB_USERNAME"))
	}

	if os.Getenv("DB_PASSWORD") == "" {
		panic(errors.New("Please provide valid DB_PASSWORD"))
	}

	if os.Getenv("DB_AUTH_DB") == "" {
		panic(errors.New("Please provide valid DB_AUTH_DB"))
	}

}
