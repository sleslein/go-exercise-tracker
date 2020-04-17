package exercise

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Log the exercise and returns a message
func Log(name string, reps int) string {
	s := fmt.Sprintf("Successfully logged: %d of %s", reps, name)
	return s
}

// Create a new exercise
func Create(name string) string {
	file, err := os.OpenFile("../exercise.txt", os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, readErr := ioutil.ReadAll(file)

	if readErr != nil {
		log.Fatal(readErr)
		os.Exit(1)
	}
	
	exercises := string(b)
	if strings.Contains(exercises, name) {
		return fmt.Sprintf("%s already exists", name)
	}

	_, writeErr := file.WriteString("\n" + name)

	if writeErr != nil {
		log.Fatal(writeErr)
	}

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
