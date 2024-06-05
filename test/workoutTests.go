package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

type WorkoutData struct {
    Duration  float64 `json:"duration"` 
    Intensity string  `json:"intensity"`
    Type      string  `json:"type"`     
}

type UserGoals struct {
    WeeklyDurationGoal float64 `json:"weeklyDurationGoal"`
    TypeGoal           string  `json:"typeGoal"`          
}

func ProcessWorkoutData(w http.ResponseWriter, req *http.Request) {
    var workout WorkoutData
    json.NewDecoder(req.Body).Decode(&workout)
    response, _ := json.Marshal(workout)
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func TrackUserGoals(w http.ResponseWriter, req *http.Request) {
    var goals UserGoals
    json.NewDecoder(req.Body).Decode(&goals)
    response, _ := json.Marshal(goals)
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func setupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/processWorkoutData", ProcessWorkoutData).Methods("POST")
    r.HandleFunc("/trackUserGoals", TrackUserGoals).Methods("POST")
    return r
}

func TestProcessWorkoutData(t *testing.T) {
    router := setupRouter()

    workoutJSON := `{"duration":45,"intensity":"high","type":"cardio"}`
    req, _ := http.NewRequest("POST", "/processWorkoutData", json.NewDecoder(json.NewDecoder(strings.NewReader(workoutJSON))))
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, 200, resp.Code)
    assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
    assert.JSONEq(t, workoutJSON, resp.Body.String())
}

func TestTrackUserGoals(t *testing.T) {
    router := setupRouter()

    goalsJSON := `{"weeklyDurationGoal":150,"typeGoal":"strength"}`
    req, _ := http.NewRequest("POST", "/trackUserGoals", json.NewDecoder(json.NewDecoder(strings.NewReader(goalsJSON))))
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, 200, resp.Code)
    assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
    assert.JSONEq(t, goalsJSON, resp.Body.String())
}

func TestMain(m *testing.M) {
    code := m.Run()
    os.Exit(code)
}