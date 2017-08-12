package tests

import "gopkg.in/mgo.v2/bson"

// Message struct for general responses
type Message struct {
	Message string `json:"message"`
}

// TokenResponse to get JWT token
type TokenResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

// User struct
type User struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Password  string        `json:"password"`
	Parent    bool          `json:"parent"`
	Email     string        `json:"email"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
}

// LimitedUser for changing a user
type LimitedUser struct {
	Parent    bool   `json:"parent"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}