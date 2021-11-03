package main

import (
	"bytes"
	"encoding/json"
	"github.com/Tutor2Tutee/T2T-GO/db"
	"github.com/Tutor2Tutee/T2T-GO/routers"
	"github.com/gin-gonic/gin"
	"github.com/tryvium-travels/memongo"
	"github.com/tryvium-travels/memongo/memongolog"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"testing"
)

func init() {
	log.SetPrefix("[TESTING] ")
	if err := os.Setenv("GO_ENV", "TEST"); err != nil {
		log.Fatalln("failed to set env")
	}
	gin.SetMode(gin.TestMode)
}

func GetMockDatabase() *db.Resource {
	opts := &memongo.Options{
		MongoVersion: "5.0.0",
		LogLevel:     memongolog.LogLevelWarn,
	}
	if runtime.GOARCH == "arm64" {
		if runtime.GOOS == "darwin" {
			// Only set the custom url as work
			opts.DownloadURL = "https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-5.0.0.tgz"
		}
	}

	memoryServer, _ := memongo.StartWithOptions(opts)
	return db.GetResource(
		"",
		"",
		memoryServer.URI(),
		"TEST",
	)
}

func TestShouldReturnNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/notapage", nil)
	w := httptest.NewRecorder()

	r := routers.GetRouter(GetMockDatabase())
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Error("Status not found should come here")
	}
}

func TestShouldGetEveryClasses(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/classes", nil)
	w := httptest.NewRecorder()

	r := routers.GetRouter(GetMockDatabase())
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("should return 200 OK")
	}
}

func TestShouldCreateASingleClass(t *testing.T) {
	//RequestBody := models.Class{Name: "HELLO", Description: "description"}
	RequestBody := gin.H{
		"name":        "HELLO",
		"description": "DESCRIPTION",
	}
	body, _ := json.Marshal(RequestBody)
	req, _ := http.NewRequest("POST", "/api/classes", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	r := routers.GetRouter(GetMockDatabase())
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Error("should return 201 CREATED")
	}
}
