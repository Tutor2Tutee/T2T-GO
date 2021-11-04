package main

import (
	"github.com/Tutor2Tutee/T2T-GO/db"
	"github.com/Tutor2Tutee/T2T-GO/routers"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	log.SetPrefix("[Main]")

	if os.Getenv("GO_ENV") != "TEST" {
		log.Println("Loading dotenv")
		if err := godotenv.Load(); err != nil {
			log.Fatalln(nil)
		}
	}

	database := db.GetResource(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_NAME"),
	)
	r := routers.GetRouter(database)
	// listen and serve on 0.0.0.0:8080
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
