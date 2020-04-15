package main

import (
	"flag"
	"fmt"
	"go-exercise-tracker/exercise"
	"os"
	"strings"
)

func main() {

	// create log sub command
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	listOption := listCommand.String("type", "exercise", "Metric {exercise|log};. (Required)")

	logCommand := flag.NewFlagSet("log", flag.ExitOnError)
	exerciseFlag := logCommand.String("name", "", "name of the exercise (required).")
	repFlag := logCommand.Int("rep", 10, "Number of Reps")

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	newExerciseName := createCommand.String("name", "", "name of the new exercise (required)")

	flag.Parse() // reads the cli args
	if len(os.Args) == 0 {
		fmt.Println("commands: list, log, create")
		fmt.Println("list [options]")
		listCommand.PrintDefaults()

		fmt.Println("log [options]")
		logCommand.PrintDefaults()

		fmt.Println("create [options]")
		createCommand.PrintDefaults()

		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCommand.Parse(os.Args[2:])
		if listCommand.Parsed() {
			listResult := ""

			switch *listOption {
			case "exercise":
				listResult = exercise.List()
			default:
				listResult = fmt.Sprintf("%s is not implemented", *listOption)
			}

			fmt.Printf(listResult)
		} else {
			listCommand.PrintDefaults()
			os.Exit(1)
		}

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
		logResult := exercise.Log(*exerciseFlag, *repFlag)
		fmt.Printf(logResult)
	}

	if createCommand.Parsed() {
		if len(strings.Trim(*newExerciseName, " ")) == 0 {
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		createResult := exercise.Create(*newExerciseName)
		fmt.Printf(createResult)
	}
}
