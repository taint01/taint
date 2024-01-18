package application

import (
	"context"
	"demo/application/model"
	"demo/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	initUser "demo/application/domains/packing/init"
)

type Server struct {
	cfg *config.Config
	lib *model.Lib
}

func NewServer(cfg *config.Config) *Server {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	fmt.Println(cfg.Database)
	opts := options.Client().ApplyURI(cfg.Database).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	db := client.Database("localdb")
	if err := db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return &Server{
		cfg: cfg,
		lib: &model.Lib{MongoDb: db},
	}
}

func (s *Server) Start() {
	s.run()
}

func (s *Server) Stop() {
	//defer func() {
	//	if err = s.lib.MongoDb.Client().Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
}

func (s *Server) run() {
	router := gin.Default()
	initHandle := initUser.NewInit(s.lib, s.cfg)
	router.GET("/get", func(c *gin.Context) {
		fmt.Println("8888")
		initHandle.Handler.Read(context.Background(), c)
	})
	router.POST("/create", func(c *gin.Context) {
		initHandle.Handler.Insert(context.Background(), c)
	})
	router.POST("/update", func(c *gin.Context) {
		initHandle.Handler.Update(context.Background(), c)
	})
	router.POST("/delete", func(c *gin.Context) {
		initHandle.Handler.Delete(context.Background(), c)
	})
	router.Run(":8080")
}
