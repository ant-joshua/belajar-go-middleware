package main

import (
	"belajar-middleware/database"
	"belajar-middleware/routers"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	database.StartDB()

	fmt.Println(os.Getenv("PORT"))

	r := routers.StartApp()
	err := r.Run(":" + os.Getenv("PORT"))

	log.Fatal(err)
}
