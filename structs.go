package main

import (
	"gopkg.in/mgo.v2/bson"
)

// User struct
type User struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Password  string        `json:"password"`
	Parent    bool          `json:"parent"`
	Email     string       `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Camp struct
type Camp struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	City     string `json:"city"`
	From     int64 `json:"from"`
	To       int64 `json:"to"`
	Location string `json:"location"`
}

// Participant struct
type Participant struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    bson.ObjectId `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Knowledge []string `json:"knowledge"`
	Interests []bson.ObjectId `json:"camp_ids"`
	Birthday  int64 `json:"birthday"`
}

// Attendant attends an camp
type Attendant struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Attended []int64 `json:"attended"`
	Laptop   bool `json:"laptop"`
	LaptopStats []struct {
		Got      bool `json:"got"`
		Received bool `json:"received"`
		ID       int  `json:"id"`
	} `json:"laptopstats"`
}
