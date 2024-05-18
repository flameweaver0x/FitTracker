package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "sync"
    "time"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

type Workout struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Reps        int    `json:"reps"`
    Sets        int    `json:"sets"`
}

type Goal struct {
    ID      string `json:"id"`
    UserID  string `json:"userId"`
    Details string `json:"details"`
}

type PersonalTrainer struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    Speciality string `json:"speciality"`
}

var (
    workouts          []Workout
    goals             []Goal
    personalTrainers  []PersonalTrainer
    workoutCache      []byte
    goalsCache        []byte
    personalTrainersCache []byte
    cacheMutex        sync.Mutex
    cacheDuration     = 10 * time.Minute
    lastCacheUpdateTime time.Time
)

func updateCache() {
    cacheMutex.Lock()
    defer cacheMutex.Unlock()

    if time.Since(lastCacheUpdateTime) > cacheDuration {
        workoutCache, _ = json.Marshal(workouts)
        goalsCache, _ = json.Marshal(goals)
        personalTrainersCache, _ = json.Marshal(personalTrainers)
        lastCacheUpdateTime = time.Now()
    }
}

func getWorkouts(w http.ResponseWriter, r *http.Request) {
    updateCache()
    sendJSONResponse(w, workoutCache)
}

func getGoals(w http.ResponseWriter, r *http.Request) {
    updateCache()
    sendJSONResponse(w, goalsCache)
}

func getPersonalTrainers(w http.ResponseWriter, r *http.Request) {
    updateCache()
    sendJSONResponse(w, personalTrainersCache)
}

func sendJSONResponse(w http.ResponseWriter, data []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
}

func initRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/api/workouts", getWorkouts).Methods("GET")
    r.HandleFunc("/api/goals", getGoals).Methods("GET")
    r.HandleFunc("/api/personal_trainers", getPersonalTrainers).Methods("GET")
    return r
}

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    seedData()

    r := initRoutes()
    port := getServerPort()
    log.Fatal(http.ListenAndServe(":"+port, r))
}

func seedData() {
    workouts = append(workouts, Workout{ID: "1", Title: "Beginner Routine", Exercises: []Exercise{{Name: "Push-ups", Description: "Standard push-ups", Reps: 10, Sets: 3}}})
    goals = append(goals, Goal{ID: "1", UserID: "42", Details: "Lose 5kg in 2 months"})
    personalTrainers = append(personalTrainers, PersonalTrainer{ID: "1", Name: "John Doe", Speciality: "Weight loss"})
}

func getServerPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    return port
}