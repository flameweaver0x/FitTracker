package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "os"
    "strings"
    "sync"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

type InMemoryCache struct {
    sync.Mutex
    items map[string]string
}

func NewInMemoryCache() *InMemoryCache {
    return &InMemoryCache{
        items: make(map[string]string),
    }
}

func (cache *InMemoryCache) GetItem(key string) (string, bool) {
    cache.Lock()
    defer cache.Unlock()
    value, found := cache.items[key]
    return value, found
}

func (cache *InMemoryCache) SetItem(key string, value string) {
    cache.Lock()
    defer cache.Unlock()
    cache.items[key] = value
}

var (
    workoutDataCache = NewInMemoryCache()
    userGoalsCache   = NewInMemoryCache()
)

type WorkoutDetails struct {
    Duration  float64 `json:"duration"`
    Intensity string  `json:"intensity"`
    Type      string  `json:"type"`
}

type GoalsDetails struct {
    WeeklyDurationGoal float64 `json:"weeklyDurationGoal"`
    TypeGoal           string  `json:"typeGoal"`
}

func HandleWorkoutDataSubmission(w http.ResponseWriter, req *http.Request) {
    var workout WorkoutDetails
    json.NewDecoder(req.Body).Decode(&workout)
    
    cacheKey := fmt.Sprintf("%f%s%s", workout.Duration, workout.Intensity, workout.Type)
    if cachedData, found := workoutDataCache.GetItem(cacheKey); found {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(cachedData))
        return
    }
    
    serializedData, _ := json.Marshal(workout)
    workoutDataCache.SetItem(cacheKey, string(serializedData))
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(serializedData)
}

func HandleUserGoalsSubmission(w http.ResponseWriter, req *http.Request) {
    var goals GoalsDetails
    json.NewDecoder(req.Body).Decode(&goals)
    
    cacheKey := fmt.Sprintf("%f%s", goals.WeeklyDurationGoal, goals.TypeGoal)
    if cachedData, found := userGoalsCache.GetItem(cacheKey); found {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(cachedData))
        return
    }
    
    serializedData, _ := json.Marshal(goals)
    userGoalsCache.SetItem(cacheKey, string(serializedData))
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(serializedSata)
}

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/submitWorkoutData", HandleWorkoutDataSubmission).Methods("POST")
    router.HandleFunc("/submitUserGoals", HandleUserGoalsSubmission).Methods("POST")
    return router
}

func TestHandleWorkoutDataSubmission(t *testing.T) {
    router := SetupRouter()

    workoutJSON := `{"duration":45,"intensity":"high","type":"cardio"}`
    req, _ := http.NewRequest("POST", "/submitWorkoutData", strings.NewReader(workoutJSON))
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, 200, resp.Code)
    assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
    assert.JSONEq(t, workoutJSON, resp.Body.String())
}

func TestHandleUserGoalsSubmission(t *testing.Br) {
    router := SetupRouter()

    goalsJSON := `{"weeklyDurationGoal":150,"typeGoal":"strength"}`
    req, _ := http.NewRequest("POST", "/submitUserGoals", strings.NewReader(goalsJSON))
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