package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type WorkoutPlan struct {
	ID      string    `json:"id"`
	Date    time.Time `json:"date"`
	Exercises []Exercise `json:"exercises"`
	Goal    string    `json:"goal"`
}

type Exercise struct {
	Name      string `json:"name"`
	Sets      int    `json:"sets"`
	Reps      int    `json:"reps"`
	Intensity string `json:"intensity"` // Low, Medium, High
}

type UserProgress struct {
	UserID        string        `json:"user_id"`
	WorkoutPlans []WorkoutPlan `json:"workout_plans"`
}

var usersProgress []UserProgress

func main() {
	// Example operations - should be replaced by actual API or CLI interactions in a real application

	// Adding a new workout log for a user
	newExercise := Exercise{Name: "Push-ups", Sets: 3, Reps: 15, Intensity: "Medium"}
	newWorkoutPlan := WorkoutPlan{ID: "1", Date: time.Now(), Exercises: []Exercise{newExercise}, Goal: "Strength"}
	AddWorkoutLog("user1", newWorkoutPlan)

	// Updating a workout log
	updatedExercise := Exercise{Name: "Push-ups", Sets: 4, Reps: 20, Intensity: "High"}
	UpdateWorkoutLog("user1", "1", []Exercise{updatedExercise})

	// Deleting a workout log
	DeleteWorkoutLog("user1", "1")

	// Suggesting a new workout plan based on user progress
	newPlan := SuggestWorkoutPlan("user1")
	fmt.Println("Suggested Plan: ", newPlan)
}

func AddWorkoutLog(userID string, plan WorkoutPlan) error {
	for i, userProgress := range usersProgress {
		if userProgress.UserID == userID {
			usersProgress[i].WorkoutPlans = append(usersProgress[i].WorkoutPlans, plan)
			return nil
		}
	}
	// User progress does not exist, create new
	newUserProgress := UserProgress{
		UserID:        userID,
		WorkoutPlans: []WorkoutPlan{plan},
	}
	usersProgress = append(usersProgress, newUserProgress)
	return saveToFile()
}

func UpdateWorkoutLog(userID, planID string, exercises []Exercise) error {
	for i, userProgress := range usersProgress {
		if userProgress.UserID == userID {
			for j, plan := range userProgress.WorkoutPlans {
				if plan.ID == planID {
					usersProgress[i].WorkoutPlans[j].Exercises = exercises
					return saveToFile()
				}
			}
		}
	}
	return errors.New("workout log not found")
}

func DeleteWorkoutLog(userID, planID string) error {
	for i, userProgress := range usersProgress {
		if userProgress.UserID == userID {
			for j, plan := range userProgress.WorkoutPlans {
				if plan.ID == planID {
					usersProgress[i].WorkoutPlans = append(usersProgress[i].WorkoutPlans[:j], usersProgress[i].WorkoutPlans[j+1:]...)
					return saveToFile()
				}
			}
		}
	}
	return errors.New("workout log not found")
}

func SuggestWorkoutPlan(userID string) *WorkoutPlan {
	// Simplified logic for suggesting a plan. Should be replaced with a more complex algorithm.
	for _, userProgress := range usersProgress {
		if userProgress.UserID == userID {
			// Assuming the last workout plan is the most recent
			if len(userProgress.WorkoutPlans) > 0 {
				lastPlan := userProgress.WorkoutPlans[len(userProgress.WorkoutPlans)-1]
				suggestedPlan := WorkoutPlan{
					ID:      fmt.Sprintf("%s-new", lastPlan.ID),
					Date:    time.Now(),
					Exercises: increaseIntensity(lastPlan.Exercises),
					Goal:    lastPlan.Goal,
				}
				return &suggestedPlan
			}
		}
	}
	return nil // No plan found for suggestion
}

func increaseIntensity(exercises []Exercise) []Exercise {
	// Simplified logic to increase exercise intensity
	for i, ex := range exercises {
		switch ex.Intensity {
		case "Low":
			exercises[i].Intensity = "Medium"
		case "Medium":
			exercises[i].Intensity = "High"
		}
	}
	return exercises
}

func saveToFile() error {
	data, err := json.Marshal(usersProgress)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("userProgress.json", data, 0644)
}

func readFromFile() error {
	file, err := ioutil.ReadFile("userProgress.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &usersProgress)
}

// In a real application, init or main should call readFromFile to load existing data.