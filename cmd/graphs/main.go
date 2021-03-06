package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graphsProject/cmd/graphs/apis"
	"log"
	"net/http"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			})
	})

	server, err := apis.GraphApi()

	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	go server.Serve()
	defer server.Close()

	fmt.Println("testfgchvjb")

	router.GET("/socket.io/", gin.WrapH(server))
	router.POST("/socket.io/", gin.WrapH(server))

	router.Run()

}
