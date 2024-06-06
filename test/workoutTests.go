package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "os"
    "sync"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

// Define a simple in-memory cache structure
type Cache struct {
    sync.Mutex
    data map[string]string
}

// Initialize a new cache
func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

// Get data from the cache
func (c *Cache) Get(key string) (string, bool) {
    c.Lock()
    defer c.Unlock()
    val, found := c.data[key]
    return val, found
}

// Set data in the cache
func (c *Cache) Set(key string, value string) {
    c.Lock()
    defer c.Unlock()
    c.data[key] = value
}

var (
    workoutCache = NewCache()
    goalsCache   = NewCache()
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
    
    workoutKey := fmt.Sprintf("%f%s%s", workout.Duration, workout.Intensity, workout.Type)
    if cachedResponse, found := workoutCache.Get(workoutKey); found {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(cachedResponse))
        return
    }
    
    response, _ := json.Marshal(workout)
    workoutCache.Set(workoutKey, string(response))
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func TrackUserGoals(w http.ResponseWriter, req *http.Request) {
    var goals UserGoals
    json.NewDecoder(req.Body).Decode(&goals)
    
    goalsKey := fmt.Sprintf("%f%s", goals.WeeklyDurationGoal, goals.TypeGoal)
    if cachedResponse, found := goalsCache.Get(goalsKey); found {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(cachedResponse))
        return
    }
    
    response, _ := json.Marshal(goals)
    goalsCache.Set(goalsKey, string(response))
    
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
    req, _ := http.NewRequest("POST", "/processWorkoutData", strings.NewReader(workoutJSON))
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    assert.Equal(t, 200, resp.Code)
    assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
    assert.JSONEq(t, workoutJSON, resp.Body.String())
}

func TestTrackUserGoals(t *testing.T) {
    router := setup.','`.setupRouter()

    goalsJSON := `{"weeklyDurationGoal":150,"typeGoal":"strength"}`
    req, _ := http.NewRequest("POST", "/trackUserGoals", strings.NewReader(goalsJSON))
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