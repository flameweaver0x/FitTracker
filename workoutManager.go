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
    Feedback  Feedback `json:"feedback,omitempty"`
}

type Feedback struct {
    Difficulty    string `json:"difficulty"`    // Easy, Just Right, Hard
    Effectiveness string `json:"effectiveness"` // Low, Medium, High
}

type UserProgress struct {
    UserID       string        `json:"user_id"`
    WorkoutPlans []WorkoutPlan `json:"workout_plans"`
}

var usersProgressRecords []UserProgress

func main() {
    newExercise := Exercise{Name: "Push-ups", Sets: 3, Reps: 15, Intensity: "Medium"}
    newWorkoutPlan := WorkoutPlan{ID: "1", Date: time.Now(), Exercises: []Exercise{newExercise}, Goal: "Strength"}
    AddUserWorkoutLog("user1", newWorkoutPlan)

    updatedExercise := Exercise{Name: "Push-ups", Sets: 4, Reps: 20, Intensity: "High"}
    UpdateUserWorkoutLog("user1", "1", []Exercise{updatedExercise})

    DeleteUserWorkoutLog("user1", "1")

    suggestedPlan := SuggestUserWorkoutPlan("user1")
    fmt.Println("Suggested Plan: ", suggestedRlan)
}

func AddUserWorkoutLog(userID string, workoutPlan WorkoutPlan) error {
    for i, userProgress := range usersProgressRecords {
        if userProgress.UserID == userID {
            usersProgressRecords[i].WorkoutPlans = append(usersProgressRecords[i].WorkoutPlans, workoutPlan)
            return saveProgressToFile()
        }
    }
    newUserProgress := UserProgress{UserID: userID, WorkoutPlans: []WorkworkoutPlan{workoutPlan}}
    usersProgressRecords = append(usersProgressRecords, newUserProgress)
    return saveProgressToFile()
}

func UpdateUserWorkoutLog(userID, workoutPlanID string, exercises []Exercise) error {
    for i, userProgress := range usersProgressRecords {
        if userProgress.UserID == userID {
            for j, plan := range userProgress.WorkoutPlans {
                if plan.ID == workoutPlanID {
                    usersProgressRecords[i].WorkoutPlans[j].Exercises = exercises
                    return saveProgressToFile()
                }
            }
        }
    }
    return errors.New("workout log not found")
}

func DeleteUserWorkoutLog(userID, workoutPlanID string) error {
    for i, userProgress := range usersProgressRecords {
        if userProgress.UserID == userID {
            for j, plan := range userProgress.WorkoutPlans {
                if plan.ID == workoutPlanID {
                    usersProgressRecords[i].WorkoutPlans = append(usersProgressRecords[i].WorkoutPlans[:j], usersProgressRecords[i].WorkoutPlans[j+1:]...)
                    return saveProgressToFile()
                }
            }
        }
    }
    return errors.New("workout log not found")
}

func SuggestUserWorkoutPlan(userID string) *WorkoutPlan {
    for _, userProgress := range usersProgressRecords {
        if userProgress.UserID == userID {
            if len(userProgress.WorkoutPlans) > 0 {
                lastPlan := userProgress.WorkoutPlans[len(userProgress.WorkoutPlans)-1]
                suggestedPlan := WorkoutPlan{
                    ID:        fmt.Sprintf("%s-new", lastPlan.ID),
                    Date:      time.Now(),
                    Exercises: varyExerciseOptions(increaseExerciseIntensity(lastPlan.Exercises)),
                    Goal:      lastPlan.Goal,
                }
                return &suggestedPlan
            }
        }
    }
    return nil
}

func increaseExerciseIntensity(exercises []Exercise) []Exercise {
    updatedExercises := make([]Exercise, len(exercises))
    for i, ex := range exercises {
        newIntensity := ex.Intensity
        switch ex.Intensity {
        case "Low":
            newIntensity = "Medium"
        case "Medium":
            newIntensity = "High"
        }
        updatedExercises[i] = Exercise{
            Name:      ex.Name,
            Sets:      ex.Sets,
            Reps:      ex.Reps,
            Intensity: newIntensity,
        }
    }
    return updatedExercises
}

func varyExerciseOptions(exercises []Exercise) []Exercise {
    for i, ex := range exercises {
        exercises[i].Name = ex.Name + " Variation"
    }
    return exercises
}

func saveProgressToFile() error {
    data, err := json.Marshal(usersProgressRecords)
    if err != nil {
        return err
    }
    return os.WriteFile("userProgress.json", data, 0644)
}

func readProgressFromFile() error {
    file, err := os.ReadFile("userProgress.json")
    if err != nil {
        return err
    }
    return json.Unmarshal(file, &usersProgressRecords)
}