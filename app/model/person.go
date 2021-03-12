package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person is the data structure that we will save and receive.
type Person struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type       string             `json:"type,omitempty" bson:"type,omitempty"`
	Attributes *Attr              `json:"attributes" bson:"attributes"`
}

// Attr is the data structure for the attributes node of a Person
type Attr struct {
	FirstName string                 `json:"firstName,omitempty" bson:"first_name,omitempty"`
	LastName  string                 `json:"lastName,omitempty" bson:"last_name,omitempty"`
	Username  string                 `json:"userName,omitempty" bson:"username,omitempty"`
	Email     string                 `json:"email,omitempty" bson:"email,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}

// NewPerson will return a Person{} instance, Person structure factory function
func NewPerson(firstName, lastName, userName, email string, data map[string]interface{}) *Person {
	a := new(Attr)
	a.Data = data
	a.Email = email
	a.FirstName = firstName
	a.LastName = lastName
	a.Username = userName

	return &Person{
		Type:       "person",
		Attributes: a,
	}
}
