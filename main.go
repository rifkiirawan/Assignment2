package main

import (
	"assignment2/database"
	"assignment2/routers"

	"fmt"
)

func main() {
	database.StartDB()
	var PORT = ":8000"
	routers.InitApiRoutes().Run(PORT)
	fmt.Println("Application is listening on port", PORT)
}
