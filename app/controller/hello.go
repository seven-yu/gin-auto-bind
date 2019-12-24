package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	params := struct {
		Name string `json:"name" form:"name"`
	}{}

	if err := c.ShouldBindQuery(&params); err != nil {
		panic(err)
	}

	fmt.Println("do something!!")
	return
}

func Create(c *gin.Context) {
	params := struct {
		Name string `json:"name" form:"name"`
		Age  int    `json:"age" form:"age"`
	}{}

	if err := c.ShouldBindJSON(&params); err != nil {
		panic(err)
	}

	fmt.Println("do something!!")
	return
}
