package main

import (
	"github.com/Camelia-hu/gee/Gee"
	"net/http"
)

func main() {
	r := Gee.New()
	r.GET("/", func(c *Gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *Gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *Gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *Gee.Context) {
		c.Json(http.StatusOK, Gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
