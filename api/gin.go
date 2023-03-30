package api

import (
	"net/http"

	"example.com/go_server/database"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Anime Review Database",
	})
	return
}

func postAnime(c *gin.Context) {
	var anime database.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.CreateAnime(&anime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
  
	c.JSON(http.StatusCreated, gin.H{
		"anime": res,
	})
	return
 }
 
 func getAnime(c *gin.Context) {
	id := c.Param("id")
	anime, err := database.ReadAnime(id)
  
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"anime": anime,
	})
	return
 }

func getAnimes(c *gin.Context) {
animes, err := database.ReadAnimes()
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"err": err,
	})
	return
}
c.JSON(http.StatusOK, gin.H{
	"animes": animes,
})
}

func putAnime(c *gin.Context) {
	var anime database.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.UpdateAnime(&anime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
  
	c.JSON(http.StatusCreated, gin.H{
		"anime": res,
	})
	return
 }
 
 func deleteAnime(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteAnime(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Anime created successfully",
	})
	return
 }

 func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home)
	r.GET("/api/v1/animes/:id", getAnime)
	r.GET("/api/v1/animes", getAnimes)
	r.POST("/api/v1/animes", postAnime)
	r.PUT("/api/v1/animes/:id", putAnime)
	r.DELETE("/api/v1/animes/:id", deleteAnime)
	return r
 }
 
 