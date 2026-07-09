package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User — структура для коллекции users
type User struct {
	ID         string `bson:"_id,omitempty"`
	Name       string `bson:"name"`
	OrderCount int    `bson:"order_count"`
}

func main() {
	// Настройка подключения
	clientOpts := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Получаем коллекцию
	collection := client.Database("test_db").Collection("users")

	// CREATE — вставка документа
	user := User{Name: "Denis", OrderCount: 5}
	insertRes, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted ID:", insertRes.InsertedID)

	// READ — выборка
	var result User
	err = collection.FindOne(ctx, bson.M{"name": "Denis"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found user:", result)

	// UPDATE — обновление
	_, err = collection.UpdateOne(ctx,
		bson.M{"name": "Denis"},
		bson.M{"$set": bson.M{"order_count": 10}},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated user")

	// DELETE — удаление
	_, err = collection.DeleteOne(ctx, bson.M{"name": "Denis"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted user")
}
