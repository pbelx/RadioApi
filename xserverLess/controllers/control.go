package controllers

import (
	radiodb "xserve/controllers/radiodb"

	"github.com/gin-gonic/gin"
)

//add station
type Xstations struct {
	Name string `json:"name" binding:"required,len=10"`
	Url  string `json:"url" binding:"required"`
}

type Dstation struct {
	Name string `json:"name" binding:"required,len=10"`
}

//home route
func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "home base",
	})

}

//get all stations radiodb "xserve/controllers/radiodb"
func Stations(c *gin.Context) {
	radiodb.Getstations(c)

}

//add new station radiodb "xserve/controllers/radiodb"
func AddStation(c *gin.Context) {

	var xx Xstations
	err := c.BindJSON(&xx)
	if err != nil {
		c.JSON(400, gin.H{
			"data": "failed",
		})
	} else {
		radiodb.Addstation(xx.Name, xx.Url, c)

	}

}

//delete station radiodb "xserve/controllers/radiodb"

func DelStation(c *gin.Context) {

	var ds Dstation
	c.BindJSON(&ds)
	radiodb.Delstation(ds.Name, c)

}

// update Url station radiodb "xserve/controllers/radiodb"
func UpdateUrl(c *gin.Context) {
	var Us Xstations
	c.BindJSON(&Us)
	radiodb.UpdateUrl(Us.Name, Us.Url, c)
}

// func Stations()
