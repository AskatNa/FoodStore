package main

import (
	"apigateway/internal/handler"
	orderPB "apigateway/proto/order"
	svc "github.com/AskatNa/apis-gen-user-service/service/frontend/client/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8082"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	orderConn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to OrderService: %v", err)
	}
	userConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to UserService: %v", err)
	}

	orderClient := orderPB.NewOrderServiceClient(orderConn)
	customerClient := svc.NewCustomerServiceClient(userConn)

	handler.InitOrderRoutes(r, orderClient)
	handler.InitUserRoutes(r, customerClient)

	log.Println("API Gateway started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("API Gateway failed: %v", err)
	}
}
