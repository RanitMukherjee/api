// tests/integration_test.go
package tests

import (
	"api/ent"
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testClient *ent.Client
var testCtx context.Context

func TestMain(m *testing.M) {
	var err error
	testClient, err = ent.Open("postgres", "postgres://hire:me@localhost:5432/hireme?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer testClient.Close()
	testCtx = context.Background()
	if err := testClient.Schema.Create(testCtx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestHabitCRUD(t *testing.T) {
	// Create
	h, err := testClient.Habit.Create().SetName("Integration").SetDescription("Test").Save(testCtx)
	if err != nil {
		t.Fatalf("Failed to create habit: %v", err)
	}

	// Read
	got, err := testClient.Habit.Get(testCtx, h.ID)
	if err != nil || got.Name != "Integration" {
		t.Fatalf("Failed to get habit or wrong data: %v", err)
	}

	// Update
	_, err = testClient.Habit.UpdateOneID(h.ID).SetName("Updated").Save(testCtx)
	if err != nil {
		t.Fatalf("Failed to update habit: %v", err)
	}

	// Delete
	err = testClient.Habit.DeleteOneID(h.ID).Exec(testCtx)
	if err != nil {
		t.Fatalf("Failed to delete habit: %v", err)
	}
}
