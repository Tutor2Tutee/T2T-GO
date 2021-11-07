package repository

import (
	"github.com/Tutor2Tutee/T2T-GO/db"
)

var Resource *db.Resource

// Collection type will ensure every collection will have Start method
type Collection interface {
	Start()
}

// SetDatabase will set database resource to our collections
func SetDatabase(database *db.Resource) {
	Resource = database
	classCollection{}.Start()
	userCollection{}.Start()
	quizCollection{}.Start()
}
