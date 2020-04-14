package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jeffdhooton/dadjokes.lol/models"
	"github.com/jinzhu/gorm"
)

// GET /jokes
// Get all jokes
func FindJokes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var jokes []models.Joke
	db.Find(&jokes)

	c.JSON(http.StatusOK, gin.H{"data": jokes})
}

// POST /jokes
// Create new joke
func CreateJoke(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.CreateJokeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	joke := models.Joke{Title: input.Title}
	db.Create(&joke)

	c.JSON(http.StatusOK, gin.H{"data": joke})
}

// GET /jokes/:id
// Find a single joke
func FindJoke(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var joke models.Joke
	if err := db.Where("id = ?", c.Param("id")).First(&joke).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Joke not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joke})
}

// PATH /jokes/:id
// Update a joke
func UpdateJoke(c *gin.Context) {
  db := c.MustGet("db").(*gorm.DB)

  // Get model if exist
  var joke models.Joke
  if err := db.Where("id = ?", c.Param("id")).First(&joke).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input models.UpdateJokeInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  db.Model(&joke).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": joke})
}

// DELETE /jokes/:id
// Delete a joke
func DeleteJoke(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var joke models.Joke
	if err := db.Where("id = ?", c.Param("id")).First(&joke).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Joke not found!"})
    return
	}

	db.Delete(&joke)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GET /random
// Get a random joke
func RandomJoke(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var jokes []models.Joke
	db.Find(&jokes)

	count := len(jokes)

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(count - 1 + 1) + 1

	var joke models.Joke
	if err := db.Where("id = ?", randIndex).First(&joke).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error fetching joke"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": joke})
}
