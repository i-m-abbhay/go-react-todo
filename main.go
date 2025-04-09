package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json: "completed"`
	Body      string `json: "body"`
}

var collection *mongo.Collection

func main() {
	fmt.Printf("hello world")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongodb local")
	collection = (client.Database("go-todo").Collection("todos"))
	app := fiber.New()
	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getTodos(c *fiber.Ctx) error {

}

// func getTodos(c *fiber.Ctx) error{

// }
// func getTodos(c *fiber.Ctx) error{

// }
// func getTodos(c *fiber.Ctx) error{

// }
