package handler

import (
    "net/http"
	"time" 
    "github.com/gin-gonic/gin"
    "github.com/coeStrayCat/golang.git/internal/db"
)

func CreateTodo (c *gin.Context) {
	var body struct {
		Title string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	type Todo struct {
		ID 				int    `json:"id"`
		Title			string `json:"title"`
		Description		string `json:"description"`
		Status 			string `json:"status"`
		CreatedAt 		time.Time `json:"created_at"`
		UpdatedAt 		time.Time `json:"updated_at"`
	}

	    var todo Todo

	err := db.DB.QueryRow(
		"INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING id, title, description, status, created_at, updated_at",
		body.Title, body.Description,
	).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}


func GetTodos(c *gin.Context) {
	type Todo struct {
		ID 				int    `json:"id"`
		Title			string `json:"title"`
		Description		string `json:"description"`
		Status 			string `json:"status"`
		CreatedAt 		time.Time `json:"created_at"`
		UpdatedAt 		time.Time `json:"updated_at"`
	}

	todos := []Todo{}
	rows, err := db.DB.Query("SELECT id, title, description, status, created_at, updated_at FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, t)
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func UpdateTodo (c *gin.Context) {
	id := c.Param("id")
	
	var body struct {
		Title string `json:"title"`
		Description string `json:"description"`
		Status string `json:"status" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	type Todo struct {
		ID 				int    `json:"id"`
		Title			string `json:"title"`
		Description		string `json:"description"`
		Status 			string `json:"status"`
		CreatedAt 		time.Time `json:"created_at"`
		UpdatedAt 		time.Time `json:"updated_at"`
	}
	
	var todo Todo
	err := db.DB.QueryRow(
		"UPDATE todos SET title = $1, description = $2, status = $3, updated_at = NOW() WHERE id = $4 RETURNING id, title, description, status, created_at, updated_at",
		body.Title, body.Description, body.Status, id,
	).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo (c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM todos WHERE id = $1", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})

}
