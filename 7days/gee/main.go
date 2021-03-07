package main

import (
	"gee"
	"net/http"
)

func main() {
	e := gee.New()

	e.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello gee</h1>")
	})

	e.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostFrom("username"),
			"password": c.PostFrom("password"),
		})
	})

	e.POST("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s is at %s\n", c.Query("name"), c.Path)
	})

	e.Run(":9999")
}
