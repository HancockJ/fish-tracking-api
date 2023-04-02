package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// catch represents data about a fish caught
type catch struct {
	ID       int     `json:"id"`
	Species  string  `json:"species"`
	Weight   float64 `json:"weight"`
	Length   float64 `json:"length"`
	Username string  `json:"username"`
}

func main() {

	router := gin.Default()
	// GET methods
	router.GET("/catches", getCatches)
	router.GET("/catches/id/:id", getCatchByID)
	router.GET("/catches/species/:species", getCatchesBySpecies)
	router.GET("/catches/username/:username", getCatchesByUsername)

	// Post methods
	router.POST("/catches", postCatches)

	router.Run("localhost:8083")
}

////////////////////
// DATA
////////////////////

// catch slice to seed catch data
var catches = []catch{
	{ID: 1, Species: "bluegill", Weight: 6.4, Length: 3.0, Username: "jackHancock"},
	{ID: 2, Species: "catfish", Weight: 28.2, Length: 12.1, Username: "jackHancock"},
	{ID: 3, Species: "carp", Weight: 25.24, Length: 20.0, Username: "jackHancock"},
	{ID: 4, Species: "bass", Weight: 17.104, Length: 9.2, Username: "jackHancock"},
	{ID: 5, Species: "trout", Weight: 14.4, Length: 7.5, Username: "jackHancock"},
	{ID: 6, Species: "bluegill", Weight: 1, Length: 2.0, Username: "jackHancock"},
}

var totalCatches = 6

////////////////////
// GET requests
////////////////////

// getCatches responds with a list of all catches as JSON.
func getCatches(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, catches)
}

// getCatchByID returns catch by ID
func getCatchByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid catch ID"})
		return
	}

	for _, i := range catches {
		if i.ID == id {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "catch not found"})
}

// getCatchBySpecies returns all catches of a certain species
func getCatchesBySpecies(c *gin.Context) {
	species := c.Param("species")
	var validCatches []catch

	for _, i := range catches {
		if i.Species == species {
			validCatches = append(validCatches, i)
		}
	}
	if len(validCatches) != 0 {
		c.IndentedJSON(http.StatusOK, validCatches)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no catches found by this species"})
	}
}

// getCatchBySpecies returns all catches of a certain species
func getCatchesByUsername(c *gin.Context) {
	username := c.Param("username")
	var validCatches []catch

	for _, i := range catches {
		if i.Username == username {
			validCatches = append(validCatches, i)
		}
	}
	if len(validCatches) != 0 {
		c.IndentedJSON(http.StatusOK, validCatches)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no catches found by this user"})
	}
}

////////////////////
// POST requests
////////////////////

// postCatches adds a catch from JSON
func postCatches(c *gin.Context) {
	var newCatch catch

	// Call BindJSON to bind the received JSON to a new catch
	if err := c.BindJSON(&newCatch); err != nil {
		return
	}

	// Format all data from new catch
	newCatch = normalizeCatchData(newCatch)

	// Increment totalCatches count and set catch ID as new total
	totalCatches += 1
	newCatch.Species = strings.ToLower(newCatch.Species)

	// Add the catch
	catches = append(catches, newCatch)
	c.IndentedJSON(http.StatusCreated, newCatch)
}

// //////////////////
// Helper functions
// //////////////////
func normalizeCatchData(newCatch catch) catch {
	newCatch.Species = strings.ToLower(newCatch.Species)
	return newCatch
}
