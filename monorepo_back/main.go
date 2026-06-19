package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello world"})
    })
    if err := r.Run(); err != nil {
        log.Fatalf("erreur au démarrage du serveur: %v", err)
    }
}