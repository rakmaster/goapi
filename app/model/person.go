package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person is the data structure that we will save and receive.
type Person struct {
	ID        primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string                 `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string                 `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Username  string                 `json:"username,omitempty" bson:"username,omitempty"`
	Email     string                 `json:"email,omitempty" bson:"email,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty" bson:"data,omitempty"` // data is a optional fields that can hold anything in key:value format.
}

// NewPerson will return a Person{} instance, Person structure factory function
func NewPerson(firstName, lastName, userName, email string, data map[string]interface{}) *Person {
	return &Person{
		Username:  userName,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Data:      data,
	}
}
