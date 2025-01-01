package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type tour struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var tours = []tour{}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // Разрешить все источники
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	connStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createToursTable(db)

	router.GET("/tours", func(c *gin.Context) {
		getTours(c, db)
	})
	router.POST("/tours", func(c *gin.Context) {
		postTours1(c, db)
	})
	router.DELETE("/tours/:id", func(c *gin.Context) {
		deleteTourByID(c, db)
	})
	router.PUT("/tours/:id", func(c *gin.Context) {
		updateTourByID(c, db)
	})

	router.Run("localhost:8084")
}

func getTours(c *gin.Context, db *sql.DB) {
	var tours []tour
	query := `SELECT id, title, description, price FROM tours`
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t tour
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tours = append(tours, t)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tours)
}

func deleteTourByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	result, err := db.Exec("DELETE FROM tours WHERE id = $1", id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка при удалении тура: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка при проверке затронутых строк: " + err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Тур не найден"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Тур удален"})
}

func updateTourByID(c *gin.Context, db *sql.DB) {
	var updatedTour tour
	if err := c.BindJSON(&updatedTour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id := c.Param("id")

	// Выполняем обновление тура
	result, err := db.Exec("UPDATE tours SET title = $1, description = $2, price = $3 WHERE id = $4", updatedTour.Title, updatedTour.Description, updatedTour.Price, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка при обновлении тура: " + err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ошибка при проверке затронутых строк: " + err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Тур не найден"})
		return
	}

	// Возвращаем обновленный тур
	c.IndentedJSON(http.StatusOK, updatedTour)
}

func postTours1(c *gin.Context, db *sql.DB) {
	var newTour tour

	if err := c.BindJSON(&newTour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO tours (title, description, price)
		VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, newTour.Title, newTour.Description, newTour.Price).Scan(&pk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert tour"})
		return
	}

	newTour.ID = pk
	c.IndentedJSON(http.StatusCreated, newTour)
}

func createToursTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS tours(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	description VARCHAR(250) NOT NULL,
	price NUMERIC(10,2) NOT NULL

	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
