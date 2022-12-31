package main

import (
	"github.com/filipecsoares/api-go-gin/database"
	"github.com/filipecsoares/api-go-gin/routes"
)

func main() {
	database.ConnectDataBase()
	routes.HandleRequests()
}
