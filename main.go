package main

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

func main() {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.Header("Content-Type", "text/html")
        c.String(http.StatusOK, `
            <div style="padding: 2rem; font-family: sans-serif;">
                <h1>Hello from Go (Gin)!</h1>
                <p>Framework test repository</p>
            </div>
        `)
    })

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":    "ok",
            "framework": "go-gin",
        })
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}