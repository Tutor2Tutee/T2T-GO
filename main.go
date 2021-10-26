package main

import (
	"github.com/Tutor2Tutee/T2T-GO/routers"
)

func main() {
	r := routers.GetRouter()

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
