package main

import (
	"demo-go/router"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	g := gin.Default()
	router.Router(g)
	err := g.Run(":8080")
	if err != nil {
		return
	}
}
