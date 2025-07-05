package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"api/ent"
	"api/ent/habit"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	// Swagger
	_ "api/docs" // This is required for swagger docs (after swag init)

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var client *ent.Client
var ctx context.Context

// @title Habit Tracker API
// @version 1.0
// @description This is an API for a habit tracker app.
// @host localhost:8080
// @BasePath /

func main() {
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

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Swagger UI endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes with named handlers
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/habits", GetHabits)
	r.POST("/habits", CreateHabit)
	r.PUT("/habits/:id", UpdateHabit)
	r.DELETE("/habits/:id", DeleteHabit)

	r.Run(":8080")
}

// @Summary Get all habits
// @Produce json
// @Success 200 {array} ent.Habit
// @Router /habits [get]
func GetHabits(c *gin.Context) {
	habits, err := client.Habit.Query().Order(ent.Desc(habit.FieldCreatedAt)).All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, habits)
}

// @Summary Create a new habit
// @Accept json
// @Produce json
// @Param habit body ent.Habit true "Habit to create"
// @Success 200 {object} ent.Habit
// @Router /habits [post]
func CreateHabit(c *gin.Context) {
	var newHabit ent.Habit
	if err := c.ShouldBindJSON(&newHabit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h, err := client.Habit.Create().
		SetName(newHabit.Name).
		SetDescription(newHabit.Description).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, h)
}

// @Summary Update a habit
// @Accept json
// @Produce json
// @Param id path int true "Habit ID"
// @Param habit body ent.Habit true "Habit to update"
// @Success 200 {object} ent.Habit
// @Router /habits/{id} [put]
func UpdateHabit(c *gin.Context) {
	id, err := toInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}
	var updatedHabit ent.Habit
	if err := c.ShouldBindJSON(&updatedHabit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h, err := client.Habit.UpdateOneID(id).
		SetName(updatedHabit.Name).
		SetDescription(updatedHabit.Description).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, h)
}

// @Summary Delete a habit
// @Produce json
// @Param id path int true "Habit ID"
// @Success 200 {object} gin.H
// @Router /habits/{id} [delete]
func DeleteHabit(c *gin.Context) {
	id, err := toInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}
	err = client.Habit.DeleteOneID(id).Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted successfully"})
}

func toInt(s string) (int, error) {
	return strconv.Atoi(s)
}
