package tests

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
	ID           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	City         string `json:"city"`
	From         int64 `json:"from"`
	To           int64 `json:"to"`
	Location     string `json:"location"`
	DocumentURLS []string `json:"document_urls"`
}

// Participant struct
type Participant struct {
	ID           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	UserID       bson.ObjectId `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Knowledge    []string `json:"knowledge"`
	Interests    []bson.ObjectId `json:"camp_ids"`
	Birthday     int64 `json:"birthday"`
	Referral     string `json:"referral"`
	Gender       string `json:"gender"`
	Diet         string `json:"diet"`
	Specialities string `json:"specialities"`
}

type rating struct {
	Grade int    `json:"grade"`
	Text  string `json:"text"`
}

// Attendant attends an camp
type Attendant struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	ParticipantID bson.ObjectId `json:"participant_id"`
	Attended      []int64 `json:"attended"`
	Laptop        bool `json:"laptop"`
	LaptopStats []struct {
		Got      bool `json:"got"`
		Received bool `json:"received"`
		ID       int  `json:"id"`
	} `json:"laptopstats"`
	LaptopPaid   bool `json:"laptop_paid"`
	FeePaid      bool `json:"paid"`
	LaptopWaived bool `json:"laptop_waive"`
	FeeWaived    bool `json:"fee_waved"`
	Documents []struct {
		Title     string `json:"title"`
		Delivered bool   `json:"delivered"`
	} `json:"laptopstats"`
	Experienced  bool `json:"experienced"`
	ProjectIdeas []string `json:"project_ideas"`
	Feedback struct {
		PrepOrganisation   rating `json:"prep_organisation"`
		DuringOrganisation rating `json:"during_organisation"`
		Feedback           string `json:"feedback"`
		Workshops          rating `json:"workshops"`
		Coaches            rating `json:"coaches"`
		Food               rating `json:"food"`
		Location           rating `json:"location"`
		Orientation        rating`json:"orientation"`
		Learned            rating`json:"learned"`
		Recommend          rating `json:"recommend"`
		EventIdeas         string `json:"event_ideas"`
	} `json:"feedback"`
}
