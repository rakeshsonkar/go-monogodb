package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rakesh/go-monogodb/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DELETEUser)
	http.ListenAndServe("localhost:9000", r)
}

// func getSession() *mgo.Session {
// 	s, err := mgo.Dial("mongodb://localhost:27017/")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }

// func getSession() *mgo.Session {
// 	// Connect to MongoDB
// 	s, err := mgo.Dial("mongodb://127.0.0.1:27017")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}

// 	// Check if the connection is successful
// 	err = s.Ping()
// 	if err != nil {
// 		log.Fatalf("Failed to ping MongoDB: %v", err)
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	return s
// }

func getSession() *mgo.Session {
	// Use the same connection string that works in Compass
	s, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	fmt.Println("Attempting to ping MongoDB...")
	// Check if the connection is successful
	err = s.Ping()
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return s
}
