package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "time"
)

type WorkoutPlan struct {
    ID        string     `json:"id"`
    Date      time.Time  `json:"date"`
    Exercises []Exercise `json:"exercises"`
    Goal      string     `json:"goal"`
}

type Exercise struct {
    Name      string   `json:"name"`
    Sets      int      `json:"sets"`
    Reps      int      `json:"reps"`
    Intensity string   `json:"intensity"` // Low, Medium, High
    Feedback  Feedback `json:"feedback,omitempty"` // include feedback on the exercise
}

type Feedback struct {
    Difficulty string `json:"difficulty"` // Easy, Just Right, Hard
    Effectiveness string `json:"effectiveness"` // Low, Medium, High
}

type UserProgress struct {
    UserID       string        `json:"user_id"`
    WorkoutPlans []WorkoutPlan `json:"workout_plans"`
}

var usersProgress []UserProgress

func main() {
    // Example operations
    newExercise := Exercise{Name: "Push-ups", Sets: 3, Reps: 15, Intensity: "Medium"}
    newWorkoutPlan := WorkoutPlan{ID: "1", Date: time.Now(), Exercises: []Exercise{newExercise}, Goal: "Strength"}
    AddWorkoutLog("user1", newWorkoutPlan)

    updatedExercise := Exercise{Name: "Push-ups", Sets: 4, Reps: 20, Intensity: "High"}
    UpdateWorkoutLog("user1", "1", []Exercise{updatedExercise})

    DeleteWorkoutLog("user1", "1")

    newPlan := SuggestWorkoutPlan("user1")
    fmt.Println("Suggested Plan: ", newPlan)

    // Let's assume reading and writing from a file works as expected, so you can persist your data.
}

func AddWorkoutLog(userID string, plan WorkoutPlan) error {
    for i, userProgress := range usersProgress {
        if userProgress.UserID == userID {
            usersProgress[i].WorkoutPlans = append(usersProgress[i].WorkoutPlans, plan)
            return saveToFile()
        }
    }
    newUserProgress := UserProgress{
        UserID:       userID,
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
    for _, userProgress := range usersProgress {
        if userProgress.UserID == userID {
            if len(userProgress.WorkoutPlans) > 0 {
                lastPlan := userProgress.WorkoutPlans[len(userProgress.WorkoutPlans)-1]
                suggestedPlan := WorkoutPlan{
                    ID:        fmt.Sprintf("%s-new", lastPlan.ID),
                    Date:      time.Now(),
                    Exercises: varyExercises(increaseIntensity(lastPlan.Exercises)), // add exercise variations
                    Goal:      lastPlan.Goal,
                }
                return &suggestedPlan
            }
        }
    }
    return nil
}

func increaseIntensity(exercises []Exercise) []Exercise {
    newExercises := make([]Exercise, len(exercises))
    for i, ex := range exercises {
        newIntensity := ex.Intensity
        switch ex.Intensity {
        case "Low":
            newIntensity = "Medium"
        case "Medium":
            newIntensity = "High"
        }
        newExercises[i] = Exercise{
            Name:      ex.Name,
            Sets:      ex.Sets,
            Reps:      ex.Reps,
            Intensity: newIntensity,
        }
    }
    return newExercises
}

// Vary exercises to include new variations
func varyExercises(exercises []Exercise) []Exercise {
    // You can expand this with real logic to select alternative exercises
    for i, ex := range exercises {
        // Add a simple variation for demonstration
        exercises[i].Name = ex.Name + " Variation"
    }
    return exercises
}

func saveToFile() error {
    data, err := json.Marshal(usersProgress)
    if err != nil {
        return err
    }
    return os.WriteFile("userProgress.json", data, 0644)
}

func readFromFile() error {
    file, err := os.ReadFile("userProgress.json")
    if err != nil {
        return err
    }
    return json.Unmarshal(file, &usersProgress)
}