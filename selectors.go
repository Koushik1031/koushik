package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name    string `bson:"name"`
	Age     int    `bson:"age"`
	Address struct {
		City    string `bson:"city"`
		State   string `bson:"state"`
		Country string `bson:"country"`
	} `bson:"address"`
}

func main() {
	// Connect to MongoDB
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Select database and collection
	db := session.DB("mydatabase")
	collection := db.C("people")

	// Find all documents with age greater than or equal to 18
	var adults []Person
	err = collection.Find(bson.M{"age": bson.M{"$gte": 18}}).All(&adults)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Adults:", adults)

	// Find all documents with age less than or equal to 18
	var children []Person
	err = collection.Find(bson.M{"age": bson.M{"$lte": 18}}).All(&children)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Children:", children)

	// Find all documents with name containing "John" case-insensitively
	var johns []Person
	err = collection.Find(bson.M{"name": bson.M{"$regex": "john", "$options": "i"}}).All(&johns)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Johns:", johns)

	// Find all documents with city equal to "New York" and state equal to "NY"
	var nyResidents []Person
	err = collection.Find(bson.M{"address.city": "New York", "address.state": "NY"}).All(&nyResidents)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NY residents:", nyResidents)

	// Find all documents with country not equal to "USA"
	var nonUSResidents []Person
	err = collection.Find(bson.M{"address.country": bson.M{"$ne": "USA"}}).All(&nonUSResidents)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Non-US residents:", nonUSResidents)
}
