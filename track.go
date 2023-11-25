package main

import (
	"errors"
	"flag"
)

func getWeightForExercise(exercise string) string {

}

func getExercises(split string) ([]string, error) {
	switch split {
	case "back":
		return []string{"", ""}, nil
	case "legs":
		return []string{"", ""}, nil
	case "chest":
		return []string{"", ""}, nil
	case "arms":
		return []string{"", ""}, nil
	case "shoulders":
		return []string{"", ""}, nil
	default:
		return nil, errors.New("workout split not defined")
	}
}

func main() {
	splitPtr := flag.String("split", "", "workout split day")
	flag.Parse()

	// Get exercises based on the flag
	exercises, err := getExercises(*splitPtr)

	if err != nil {
		// do something bad
	}

	for _, exercise := range exercises {

		// get the weight
		weight := getWeightForExercise(exercise)

		// execute the queries
	}

}
