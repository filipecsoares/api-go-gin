package main

import (
	"github.com/filipecsoares/api-go-gin/models"
	"github.com/filipecsoares/api-go-gin/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "Filipe", Id: 1, Email: "filipe@email.com"},
		{Name: "Tumo", Id: 2, Email: "tumo@email.com"},
		{Name: "Tuminho", Id: 3, Email: "tuminho@email.com"},
		{Name: "Mae deles", Id: 4, Email: "maedeles@email.com"},
	}
	routes.HandleRequests()
}
