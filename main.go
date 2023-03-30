package main

import (
	
	api "example.com/go_server/api"
	"example.com/go_server/database"
)

func init() {
	database.NewPostgreSQLClient()
}


func main() {
	r := api.SetupRouter()

	r.Run(":8080")
}	



