package main

import (
	"fmt"
	"log"
	"microservice_grpc_wishlist/db"
	"microservice_grpc_wishlist/pb/wishlist"
	"microservice_grpc_wishlist/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("failed connect database: %v", err)
	}

	grpcServer := grpc.NewServer()
	wishlistService := services.WishlistService{
		DB: db,
	}

	wishlist.RegisterWishlistServiceServer(grpcServer, &wishlistService)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen on port :8080 %v", err)
	}

	fmt.Println("server running on port :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server grpc server %v", err)
	}
}
