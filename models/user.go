package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"id" bson:"name"`
	Gender string        `json:"id" bson:"gender"`
	Age    int           `json:"id" bson:"age"`
}
