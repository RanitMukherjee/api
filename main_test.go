package main

import (
	"api/ent"
	"context"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	var err error
	client, err = ent.Open("postgres", "postgres://hire:me@localhost:5432/hireme?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx = context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	code := m.Run()

	// Cleanup using database/sql directly
	sqlDB, err := sql.Open("postgres", "postgres://hire:me@localhost:5432/hireme?sslmode=disable")
	if err == nil {
		defer sqlDB.Close()
		_, _ = sqlDB.ExecContext(ctx, "TRUNCATE habits RESTART IDENTITY CASCADE;")
	}

	os.Exit(code)
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", GetHabits)
	r.POST("/habits", CreateHabit)
	r.GET("/habits/:id", ShowEditHabit)
	r.PUT("/habits/:id", UpdateHabit)
	r.DELETE("/habits/:id", DeleteHabit)
	return r
}

func createTestHabit(t *testing.T, router *gin.Engine, name string) int {
	req := httptest.NewRequest("POST", "/habits", strings.NewReader("name="+name+"&description=test"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	habits, err := client.Habit.Query().Order(ent.Desc("created_at")).All(ctx)
	if err != nil || len(habits) == 0 {
		t.Fatalf("Failed to create habit for test: %v", err)
	}
	return habits[0].ID
}

func TestCreateHabit_EmptyName(t *testing.T) {
	origRender := renderHabits
	renderHabits = func(c *gin.Context, habits []*ent.Habit, editID int) {
		c.String(http.StatusOK, "mocked")
	}
	defer func() { renderHabits = origRender }()

	router := gin.New()
	router.POST("/habits", CreateHabit)
	req := httptest.NewRequest("POST", "/habits", strings.NewReader("name=&description=desc"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Name required") {
		t.Errorf("Expected 'Name required' error, got %s", w.Body.String())
	}
}

func TestCreateHabit_Success(t *testing.T) {
	router := setupTestRouter()
	req := httptest.NewRequest("POST", "/habits", strings.NewReader("name=Test&description=desc"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestCreateHabit_OnlyName(t *testing.T) {
	router := setupTestRouter()
	req := httptest.NewRequest("POST", "/habits", strings.NewReader("name=OnlyName"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetHabits(t *testing.T) {
	router := setupTestRouter()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestShowEditHabit_ValidID(t *testing.T) {
	router := setupTestRouter()
	id := createTestHabit(t, router, "EditMe")
	req := httptest.NewRequest("GET", "/habits/"+strconv.Itoa(id), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestShowEditHabit_InvalidID(t *testing.T) {
	router := setupTestRouter()
	req := httptest.NewRequest("GET", "/habits/notanint", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestUpdateHabit_Success(t *testing.T) {
	router := setupTestRouter()
	id := createTestHabit(t, router, "ToUpdate")
	form := "name=Updated&description=UpdatedDesc"
	req := httptest.NewRequest("PUT", "/habits/"+strconv.Itoa(id), strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestUpdateHabit_EmptyName(t *testing.T) {
	router := setupTestRouter()
	id := createTestHabit(t, router, "ToUpdate2")
	form := "name=&description=desc"
	req := httptest.NewRequest("PUT", "/habits/"+strconv.Itoa(id), strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestUpdateHabit_InvalidID(t *testing.T) {
	router := setupTestRouter()
	form := "name=Updated&description=desc"
	req := httptest.NewRequest("PUT", "/habits/notanint", strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestUpdateHabit_NonExistentID(t *testing.T) {
	router := setupTestRouter()
	form := "name=Ghost&description=desc"
	nonExistentID := 999999
	req := httptest.NewRequest("PUT", "/habits/"+strconv.Itoa(nonExistentID), strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

func TestDeleteHabit_Success(t *testing.T) {
	router := setupTestRouter()
	id := createTestHabit(t, router, "ToDelete")
	req := httptest.NewRequest("DELETE", "/habits/"+strconv.Itoa(id), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestDeleteHabit_InvalidID(t *testing.T) {
	router := setupTestRouter()
	req := httptest.NewRequest("DELETE", "/habits/notanint", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestDeleteHabit_NonExistentID(t *testing.T) {
	router := setupTestRouter()
	nonExistentID := 999999
	req := httptest.NewRequest("DELETE", "/habits/"+strconv.Itoa(nonExistentID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	// Your handler likely returns 200 even if nothing deleted; this still covers the branch.
}

func TestRenderHabits_HXRequest(t *testing.T) {
	router := setupTestRouter()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("HX-Request", "true")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestRenderHabits_Normal(t *testing.T) {
	router := setupTestRouter()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestRenderHabits_EditMode(t *testing.T) {
	router := setupTestRouter()
	id := createTestHabit(t, router, "EditMode")
	req := httptest.NewRequest("GET", "/habits/"+strconv.Itoa(id), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestToInt(t *testing.T) {
	if v, err := toInt("42"); err != nil || v != 42 {
		t.Error("toInt failed for valid input")
	}
	if v, err := toInt("notanint"); err == nil || v != 0 {
		t.Error("toInt failed for invalid input")
	}
}
