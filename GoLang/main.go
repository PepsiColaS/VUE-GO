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
	ID          string  `json:"id"`
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
	// router.GET("/tours/:id", getTourByID)
	router.POST("/tours", func(c *gin.Context) {
		postTours1(c, db)
	})
	// router.DELETE("/tours/:id", deleteTourByID)
	// router.PUT("/tours/:id", updateTourByID)

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
	defer rows.Close() // Закрываем rows после завершения работы с ними

	for rows.Next() {
		var t tour
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tours = append(tours, t) // Добавляем тур в срез
	}

	// Проверяем на наличие ошибок после завершения обработки
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tours)
}

func getTourByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range tours {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteTourByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range tours {
		if a.ID == id {
			// Удаляем альбом из среза
			tours = append(tours[:i], tours[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateTourByID(c *gin.Context) {
	id := c.Param("id")
	var updatedTour tour

	// Пробуем привязать JSON из запроса к структуре Album
	if err := c.BindJSON(&updatedTour); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// Ищем альбом по ID
	for i, a := range tours {
		if a.ID == id {
			// Обновляем поля альбома
			tours[i].Title = updatedTour.Title
			tours[i].Description = updatedTour.Description
			tours[i].Price = updatedTour.Price
			c.IndentedJSON(http.StatusOK, tours[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func createToursTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS tours(
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	description VARCHAR(250) NOT NULL,
	price NUMERIC(10,2) NOT NULL,
	img VARCHAR(255)
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
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
	c.IndentedJSON(http.StatusCreated, newTour)
}
