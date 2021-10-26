package main

import (
	"github.com/Tutor2Tutee/T2T-GO/db"
	"github.com/Tutor2Tutee/T2T-GO/routers"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	log.SetPrefix("[Main]")
	r := routers.GetRouter()
	db.MongoConn()
	// listen and serve on 0.0.0.0:8080
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
