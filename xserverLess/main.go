package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	h "xserve/handlers"
)

//
func main() {
	fmt.Println("server 1")

	r := gin.Default()
	r.GET("/", h.HomeRoute)             //home route
	r.GET("/stations", h.StationsRoute) //get all station
	r.POST("/addstation", h.AddStation) //add new stations Post name,url json
	r.POST("/delstation", h.DelStation) //delete stations  Post name json
	r.POST("/update", h.UpdateUrl)      //update stations Post json name,url
	r.Run(":9000")
}
