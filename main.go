package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getWeightForExercise(exercise string) int {
	fmt.Printf("Please enter your GVT weight this week for %s: ", exercise)
	var weightString string

	_, err := fmt.Scanln(&weightString)
	if err != nil {
		log.Fatal(err)
	}

	weight, err := strconv.Atoi(weightString)
	if err != nil {
		log.Fatal(err)
	}

	return weight
}

func getExercises(split string) ([]string, error) {
	switch split {
	case "back":
		return []string{"barbell-row", "weighted-pullup"}, nil
	case "legs":
		return []string{"deadlift", "front-squat"}, nil
	case "chest":
		return []string{"incline-dumbbell-press", "bench-press"}, nil
	case "arms":
		return []string{"barbell-curl", "weighted-dip"}, nil
	case "shoulders":
		return []string{"arnold-press", "overhead-press"}, nil
	default:
		return nil, errors.New("workout split not defined")
	}
}

func main() {

	now := time.Now().Format("2006-01-02 15:04:05")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/workout_tracker")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	splitPtr := flag.String("split", "", "workout split day")
	flag.Parse()

	// Get exercises based on the flag
	exercises, err := getExercises(*splitPtr)
	if err != nil {
		log.Fatal(err)
	}

	for _, exercise := range exercises {
		// get the weight
		weight := getWeightForExercise(exercise)

		// insert into the table
		query := fmt.Sprintf("INSERT INTO exercise_progress (exercise, date_recorded, weight_in_lbs) VALUES ('%s', '%s', %d);", exercise, now, weight)
		fmt.Printf("%s\n\n", query)
		_, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
	}

}
