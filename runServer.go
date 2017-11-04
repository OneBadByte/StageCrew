package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	//creates the 
	router := gin.Default()
	port := ":8080"

	//does all the http requests here

	//sets the root of the directory to access
	router.Static("/css", "./web/css")
	router.Static("/js", "./web/js")
	router.Static("/images", "./web/images")
	router.Static("/json", "./web/json")
	router.Static("/html", "./web/html")
	router.StaticFile("/", "./web/html/index.html")
	//starts web server
	router.Run(port)
}

//http functions

