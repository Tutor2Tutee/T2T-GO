package main

import (
	"github.com/Tutor2Tutee/T2T-GO/db"
	"github.com/Tutor2Tutee/T2T-GO/routers"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	log.SetPrefix("[Main]")
	r := routers.GetRouter()
	client := db.MongoConn(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_URL"))
	// listen and serve on 0.0.0.0:8080
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
