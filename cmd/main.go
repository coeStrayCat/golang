package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/joho/godotenv"
    "github.com/coeStrayCat/golang.git/internal/db"
    "github.com/coeStrayCat/golang.git/internal/handler"
)

func main() {
    godotenv.Load()
    db.Connect()

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    v1 := r.Group("/api/v1")
    {
        v1.GET("/users", handler.ListUsers)
        v1.GET("/users/:id", handler.GetUser)
        v1.POST("/users", handler.CreateUser)
        v1.DELETE("/users/:id", handler.DeleteUser)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(r.Run(":" + port))
}