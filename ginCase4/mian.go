package main

import (
	"github.com/gin-gonic/gin"
	"golang_Gin/ginCase4/routers"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/**/*")
	r.Static("/static", "./static")
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
