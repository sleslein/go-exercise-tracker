package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// create log sub command
	logCommand := flag.NewFlagSet("log", flag.ExitOnError)
	exerciseFlag := logCommand.String("name", "", "name of the exercise (required).")
	repFlag := logCommand.Int("rep", 10, "Number of Reps")

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	newExerciseName := createCommand.String("Name", "", "name of the new exercise (required)")

	flag.Parse() // reads the cli args
	if len(os.Args) < 2 {
		fmt.Println("commands: log, create")
		fmt.Println("log [options]")
		logCommand.PrintDefaults()
		fmt.Println("create [options]")
		createCommand.PrintDefaults()

		os.Exit(1)
	}

	switch os.Args[1] {
	case "log":
		logCommand.Parse(os.Args[2:])

	case "create":
		createCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if logCommand.Parsed() {
		if len(strings.Trim(*exerciseFlag, " ")) == 0 {
			logCommand.PrintDefaults()
			os.Exit(1)
		}

		fmt.Printf("successfully logged: %d of %s", *repFlag, *exerciseFlag)
	}

	if createCommand.Parsed() {
		if len(strings.Trim(*newExerciseName, " ")) == 0 {
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		fmt.Printf("successfully created: %s", *newExerciseName)
	}

}
