package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e1 := e.Group("api/")
	e1.GET("/v2/items", post.getQiita)

	e.Start(":9000")
}
