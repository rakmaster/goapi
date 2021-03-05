package config

import (
	"fmt"
	"os"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestConfig(t *testing.T) {
	os.Setenv("server_host", "localhost:8080")
	os.Setenv("mongo_user", "root")
	os.Setenv("mongo_password", "password")
	os.Setenv("mongo_host", "mongodb")
	os.Setenv("mongo_port", "27017")

	config := NewConfig()

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		"root",
		"password",
		"mongodb",
		"27017",
	)

	if config.MongoURI() != mongoURI {
		t.Fatalf("%s there is an problem in mongo connection uri generator. %s != %s", failed, config.MongoURI(), mongoURI)
	}
	t.Logf("%s Testing MongoDB connection uri generator is successful", succeed)
}
