package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Drink milk", Completed: true},
	{ID: "3", Item: "Take a walk", Completed: false},
}

// gett todos
func getTodos(context *gin.Context) {
	// context holds the data coming from the http request
	context.IndentedJSON(http.StatusOK, todos)
}

// create a todo
func createTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// return the index, the element, and an error
func getTodoById(id string) (int, *todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return i, &todos[i], nil
		}
	}

	return 0, nil, errors.New("Todo not found")
}

// get single todo
func getTodo(context *gin.Context) {
	// extract the id parameter from the url
	id := context.Param("id")
	_, todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	_, todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)

}

// delete todo
func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	index, _, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}

	todos = append(todos[:index], todos[index+1:]...)
	context.IndentedJSON(http.StatusOK, gin.H{
		"message": "resource deleted successfully",
	})
}

func main() {
	// create a gin router
	router := gin.Default()

	// create an endpoint
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.DELETE("/todos/:id", deleteTodo)

	// run the router
	router.Run(":9999")

}
