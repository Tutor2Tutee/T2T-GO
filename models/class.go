package models

import "time"

type Class struct {
	// Class fields
	Name        string    `json:"name"`         // Name of this class, required, unique
	Description string    `json:"description"`  // Description of this class, optional
	ClassType   string    `json:"class_type"`   // ClassType of online, offline, q&a, video
	MaxCapacity int       `json:"max_capacity"` // MaxCapacity of class, required, max capacity should be 50
	State       string    `json:"state"`        // State of this class, default to be Pending, have a state of pending, start, end
	Teacher     User      `json:"teacher"`      // Teacher creator of class
	Created_At  time.Time `json:"created_at"`
	Listener    []User    `json:"listener"` // Listener who listens this class
}
