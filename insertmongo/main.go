package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Age    float64 `json:"age"`
	Social string  `json:"social"`
}

func ServeHTTP() string {
	url := "http://api:8080/api/v1/getstds"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	return responseString
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:example@mongo:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	var dbName string = "bgome"
	var dbCol string = "bcol"
	collection := client.Database(dbName).Collection(dbCol)
	outJSON := ServeHTTP()

	var data map[string]interface{}
	json.Unmarshal([]byte(outJSON), &data)

	var _name, _social, _type string
	var _age float64
	for _, bar := range data["users"].([]interface{}) {
		// parse map
		_name = bar.(map[string]interface{})["name"].(string)
		_social = bar.(map[string]interface{})["social"].(string)
		_type = bar.(map[string]interface{})["type"].(string)
		_age = bar.(map[string]interface{})["Age"].(float64)
		// debug -- fmt.Println("hihi", _name, _type, _age, _social)

		data := User{_name, _type, _age, _social}
		insertResult, err := collection.InsertOne(context.TODO(), data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted document: ", insertResult.InsertedID)
	}
}
