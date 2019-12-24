package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetParams struct {
	Name string `json:"name" form:"name"`
}

func AutoBindGet(c *gin.Context, params *GetParams) {
	fmt.Println("do something!!")
	return
}
