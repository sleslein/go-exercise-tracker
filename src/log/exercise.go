package exercise

import "fmt"

// Log the exercise and returns a message
func Log(name string, reps int) string {
	s := fmt.Sprintf("Successfully logged: %d of %s", reps, name)
	return s
}

// Create a new exercise
func Create(name string) string {
	return fmt.Sprintf("successfully created: %s", name)
}
