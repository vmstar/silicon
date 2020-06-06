package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("html/*")

	router.GET("/string/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.Writer.WriteString(fmt.Sprintf("Hello %s", name))
	})

	router.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.html", "flysnow_org")
	})

	s := &http.Server{
		Addr:           ":8081",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("ListenAndServe failed. error=%v", err)
	}
}