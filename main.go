package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jairhdev/go-api-contact/config"
	"github.com/jairhdev/go-api-contact/controller"
	"github.com/jairhdev/go-api-contact/external/database"
	"github.com/jairhdev/go-api-contact/external/messaging"
)

func init() {
	const fileEnv string = "environment.properties"
	env, err := readEnvironment(fileEnv)
	if err != nil {
		panic(err)
	}

	// Start config
	if err := config.NewConfig(env); err != nil {
		panic(err)
	}

	// Start database
	var db = database.NewService(database.NewDatabase())
	if err := db.NewConnect(); err != nil {
		panic(err)
	}

	// Start messaging
	var mq = messaging.NewService(messaging.NewMessaging())
	if err := mq.NewConnect(); err != nil {
		panic(err)
	}
}

func main() {
	defer closeResources()

	controller.StartContactRoutes()

	port := os.Getenv("HOST_PORT")
	fmt.Printf("API running on port: %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func readEnvironment(fileEnv string) (string, error) {
	_, err := os.Stat(fileEnv)
	if err != nil {
		if os.IsNotExist(err) {
			return "", err
		}
	}
	result, err := os.ReadFile(fileEnv)
	return strings.TrimSpace(string(result)), err
}

func closeResources() {
	database.NewDatabase().CloseConn()
	messaging.NewMessaging().CloseConn()
	messaging.NewMessaging().CloseLogQueue()
}
