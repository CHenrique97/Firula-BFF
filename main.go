package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/UserFinder/models"

	"google.golang.org/grpc"
)

 

func main() {


	fmt.Println("Hello World")
		// Set up a connection to the gRPC server
		conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to the server: %v", err)
		}
		defer conn.Close()
	
		// Create a client instance for the UserService
		client := pb.NewUserServiceClient(conn)
	
		// Example: Create a user
		req := &pb.UserRequest{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}
	
		
		response, err := client.CreateUser(context.Background(), req)
		if err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
	
		log.Printf("User created successfully. ID: %s", response.Id)
	
		// Example: Get a user
		req = &pb.UserRequest{
			Email:    "john@example.com",
			Password: "password123",
		}
	
		response, err = client.GetUser(context.Background(), req)
		if err != nil {
			log.Fatalf("Failed to get user: %v", err)
		}
	
		log.Printf("User retrieved successfully. ID: %s, Name: %s", response.Id, response.Name)
	
}
