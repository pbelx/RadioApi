package handlers

// import "github.com/gin-gonic/gin"

import (
	cr "xserve/controllers"

	"github.com/gin-gonic/gin"
)

//Get home /
func HomeRoute(c *gin.Context) {
	cr.Home(c)
}

//Get stations /stations
func StationsRoute(c *gin.Context) {
	cr.Stations(c)
}

//Post stations /addstation
func AddStation(c *gin.Context) {
	cr.AddStation(c)
}

//delete stations /delstation
func DelStation(c *gin.Context) {
	cr.DelStation(c)
}

// update url
func UpdateUrl(c *gin.Context) {
	cr.UpdateUrl(c)
}
