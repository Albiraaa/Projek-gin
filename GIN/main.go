package main

import (
	"database/sql"
	"os"
	"projek/controllers"
	"projek/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	// Pakai DATABASE_URL langsung (lebih aman)
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is empty — Railway didn't load DB variables")
	}

	// Tambah sslmode=require (WAJIB di Railway Postgres)
	dsn = dsn + "?sslmode=require"

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

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

	// Default PORT (Railway kasih env PORT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback untuk local development
	}

	router.Run(":" + port)
}
