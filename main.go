package main

import (
	"log"

	"github.com/adityasinghtads/api_go/database"
	"github.com/adityasinghtads/api_go/routers"
)

func main() {
	database.Setup()                    // establishing the database connection
	engine := routers.Router()          // get the customized engine
	err := engine.Run("127.0.0.1:8080") // start the engine
	if err != nil {
		log.Fatal(err)
	}
}
