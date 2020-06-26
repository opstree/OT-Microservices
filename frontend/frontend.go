package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "templates/static")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
	})
    router.GET("/list", func(c *gin.Context) {
        c.HTML(http.StatusOK, "list.html", nil)
	})
	router.Run(":8000")
}
