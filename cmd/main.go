package main

import (
    "log"
	"fmt"
    "os"
    // "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/joho/godotenv"
    "github.com/coeStrayCat/golang.git/internal/db"
    "github.com/coeStrayCat/golang.git/internal/handler"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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

		v1.POST("/todo", handler.CreateTodo)
		v1.GET("/todos", handler.GetTodos)
		v1.DELETE("/todos/:id", handler.DeleteTodo)
		v1.PUT("/todos/:id", handler.UpdateTodo)
    }

    // http.HandleFunc("/", func(w http. ResponseWriter, r*http.Request) {
	// w.Write([]byte("Hello, World!"))
	// })
	// fmt.Println("Server is running on http://localhost:3000")
	// http.ListenAndServe(":3000", nil)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

	str := myForLoop(5)
	fmt.Println(str)

	sumResult := sum(10, 5)
	fmt.Println("Sum:", sumResult)

	subtractResult := subtract(10, 5)
	fmt.Println("Subtract:", subtractResult)
	
	multiplyResult := multiply(10, 5)
	fmt.Println("Multiply:", multiplyResult)
	
	divideResult, err := divide(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide:", divideResult)
	}

    log.Fatal(r.Run(":" + port))

}

func myForLoop(count int) string {
    for i := 0; i < count; i++ {
        fmt.Println("My count is:", i)  
    }
    return "Loop completed"
}

func sum(a, b int) int{
	return a + b
}

func subtract(a, b int) int{
	return a - b
}

func multiply(a, b int) int{
	return a * b
}

func divide(a, b int) (int, error){
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}




	