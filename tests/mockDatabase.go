package tests

import (
	"github.com/Tutor2Tutee/T2T-GO/db"
	"github.com/tryvium-travels/memongo"
	"github.com/tryvium-travels/memongo/memongolog"
	"runtime"
)

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
