package main

import (
	"database/sql"
	"fmt"
	"os"
	"projek/controllers"
	"projek/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGNAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	router := gin.Default()
	router.GET("/bioskop", controllers.GetAllbioskop)
	router.GET("/bioskop/:id", controllers.GetBioskopByID)
	router.POST("/bioskop", controllers.InsertBioskop)
	router.PUT("/bioskop/:id", controllers.UpdateBioskop)
	router.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	router.Run(":" + os.Getenv("PORT"))
}
