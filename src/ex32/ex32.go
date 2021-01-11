package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", Login)
	log.Fatalln(router.Run(":8080"))
}