package controllers

import (
	"net/http"
	"projek/database"
	"projek/models"
	"projek/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllbioskop(c *gin.Context) {
	var (
		result gin.H
	)

	person, err := repository.GetAllbioskop(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": person,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBioskop(c *gin.Context) {
	var bioskop models.Bioskop

	err := c.BindJSON(&bioskop)
	if err != nil {
		panic(err)
	}

	err = repository.InsertBioskop(database.DbConnection, bioskop)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, bioskop)
}

func UpdateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&bioskop)
	if err != nil {
		panic(err)
	}

	bioskop.ID = id

	err = repository.UpdateBioskop(database.DbConnection, bioskop)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, bioskop)
}

func DeleteBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	id, _ := strconv.Atoi(c.Param("id"))

	bioskop.ID = id
	err := repository.DeleteBioskop(database.DbConnection, bioskop)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, bioskop)
}

func GetBioskopByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	bioskop, err := repository.GetBioskopByID(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}
