package handler

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/coeStrayCat/golang.git/internal/db"

)

func ListUsers(c *gin.Context) {
    rows, err := db.DB.Query("SELECT id, name, email, created_at FROM users ORDER BY id")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    type User struct {
        ID        int    `json:"id"`
        Name      string `json:"name"`
        Email     string `json:"email"`
        CreatedAt string `json:"created_at"`
    }

    users := []User{}
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
        users = append(users, u)
    }

    c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
    id := c.Param("id")

    type User struct {
        ID        int    `json:"id"`
        Name      string `json:"name"`
        Email     string `json:"email"`
        CreatedAt string `json:"created_at"`
    }

    var u User
    err := db.DB.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", id).
        Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": u})
}

func CreateUser(c *gin.Context) {
    var body struct {
        Name  string `json:"name" binding:"required"`
        Email string `json:"email" binding:"required"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    type User struct {
        ID        int    `json:"id"`
        Name      string `json:"name"`
        Email     string `json:"email"`
        CreatedAt string `json:"created_at"`
    }

    var u User
    err := db.DB.QueryRow(
        "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at",
        body.Name, body.Email,
    ).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": u})
}

func DeleteUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    result, err := db.DB.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}