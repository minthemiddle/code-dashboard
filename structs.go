package main

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password"`
}
