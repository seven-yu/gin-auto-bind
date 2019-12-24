package main

import (
	"github.com/gin-gonic/gin"
	"github.com/seven-yu/gin-auto-bind/app/router"
)

func main() {
	e:=gin.Default()
	router.RouteV1(e)

	err := e.Run(":8080")
	if err != nil {
		panic(err)
	}
}


