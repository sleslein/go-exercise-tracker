package exercise

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Log the exercise and returns a message
func Log(name string, reps int) string {
	s := fmt.Sprintf("Successfully logged: %d of %s", reps, name)
	return s
}

// Create a new exercise
func Create(name string) string {
	return fmt.Sprintf("successfully created: %s", name)
}

// List all exercises
func List() string {
	file, err := os.Open("../exercise.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	return string(b)
}
