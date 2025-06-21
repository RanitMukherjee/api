package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"api/ent"
	"api/ent/habit"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	// Swagger
	_ "api/docs" // This is required for swagger docs (after swag init)

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Habit Tracker API
// @version 1.0
// @description This is an API for a habit tracker app.
// @host localhost:8080
// @BasePath /

func main() {
	client, err := ent.Open("postgres", "postgres://hire:me@localhost:5432/hireme?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Swagger UI endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Helper to render either full page or just the table body for htmx
	renderHabits := func(c *gin.Context, habits []*ent.Habit, editID int) {
		isHX := c.GetHeader("HX-Request") == "true"
		data := gin.H{"habits": habits, "editID": editID}
		if isHX {
			c.HTML(http.StatusOK, "body.html", data)
		} else {
			c.HTML(http.StatusOK, "index.html", data)
		}
	}

	// @Summary Get all habits
	// @Produce html
	// @Success 200 {string} string "HTML page"
	// @Router / [get]
	r.GET("/", func(c *gin.Context) {
		habits, _ := client.Habit.Query().Order(ent.Desc(habit.FieldCreatedAt)).All(ctx)
		renderHabits(c, habits, 0)
	})

	// @Summary Create a new habit
	// @Accept application/x-www-form-urlencoded
	// @Produce html
	// @Param name formData string true "Name"
	// @Param description formData string false "Description"
	// @Success 200 {string} string "HTML page"
	// @Router /habits [post]
	r.POST("/habits", func(c *gin.Context) {
		name := c.PostForm("name")
		desc := c.PostForm("description")
		if strings.TrimSpace(name) == "" {
			c.String(http.StatusBadRequest, "Name required")
			return
		}
		_, _ = client.Habit.Create().
			SetName(name).
			SetDescription(desc).
			Save(ctx)
		habits, _ := client.Habit.Query().Order(ent.Desc(habit.FieldCreatedAt)).All(ctx)
		renderHabits(c, habits, 0)
	})

	// @Summary Show edit form for a habit
	// @Produce html
	// @Param id path int true "Habit ID"
	// @Success 200 {string} string "HTML page"
	// @Router /habits/{id} [get]
	r.GET("/habits/:id", func(c *gin.Context) {
		id := toInt(c.Param("id"))
		habits, _ := client.Habit.Query().Order(ent.Desc(habit.FieldCreatedAt)).All(ctx)
		renderHabits(c, habits, id)
	})

	// @Summary Update a habit
	// @Accept application/x-www-form-urlencoded
	// @Produce html
	// @Param id path int true "Habit ID"
	// @Param name formData string true "Name"
	// @Param description formData string false "Description"
	// @Success 200 {string} string "HTML page"
	// @Router /habits/{id} [put]
	r.PUT("/habits/:id", func(c *gin.Context) {
		id := toInt(c.Param("id"))
		name := c.PostForm("name")
		desc := c.PostForm("description")
		if strings.TrimSpace(name) == "" {
			c.String(http.StatusBadRequest, "Name required")
			return
		}
		_, err := client.Habit.UpdateOneID(id).
			SetName(name).
			SetDescription(desc).
			Save(ctx)
		if err != nil {
			c.String(http.StatusInternalServerError, "Update failed")
			return
		}
		habits, _ := client.Habit.Query().Order(ent.Desc(habit.FieldCreatedAt)).All(ctx)
		renderHabits(c, habits, 0)
	})

	// @Summary Delete a habit
	// @Produce html
	// @Param id path int true "Habit ID"
	// @Success 200 {string} string "HTML page"
	// @Router /habits/{id} [delete]
	r.DELETE("/habits/:id", func(c *gin.Context) {
		id := toInt(c.Param("id"))
		client.Habit.DeleteOneID(id).Exec(ctx)
		habits, _ := client.Habit.Query().Order(ent.Desc(habit.FieldCreatedAt)).All(ctx)
		renderHabits(c, habits, 0)
	})

	r.Run(":8080")
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
