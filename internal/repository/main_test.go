package repository

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/tarkovskynik/Golang-ninja-project-2/pkg/database/mongodb"
)

var TestUserRepository UserInterface
var Ctx context.Context

func TestMain(m *testing.M) {
	Ctx, _ = context.WithTimeout(context.Background(), time.Second*5)

	// TODO fix to dynamic config params
	cfg := mongodb.NewConfig("mongodb://localhost:27017", "admin", "qwerty123")

	client, err := mongodb.NewClient(Ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	TestUserRepository = NewUser(client.Database("users"), "users")

	m.Run()
}
