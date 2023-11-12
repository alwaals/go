package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ToDo struct{
	ID string `json:"Id"`
	Name string `json:"Title"`
	Completed bool `json:"Completed"`
}
var todos = []ToDo{
	{ID:"101",Name: "Priya",Completed: false},
	{ID:"102",Name: "Nihira",Completed: false},
	{ID:"103",Name: "Swetha",Completed: false},
	{ID:"104",Name: "Ruchi",Completed: true},
}
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK,todos)
}
func addTodos(c *gin.Context){
	var todo ToDo
	err := c.BindJSON(&todo)
	if err!=nil{
		c.IndentedJSON(http.StatusInternalServerError,err)
	}
	todos = append(todos, todo)
	c.IndentedJSON(http.StatusCreated,todo)
}
func getById(c *gin.Context){
	var t ToDo
	err:=c.BindJSON(&t)
	if err!=nil{
		return
	}
	fmt.Println("Id received:",t.ID)
	for _,todo := range todos{
		if t.ID==todo.ID{
			c.IndentedJSON(http.StatusOK,todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,t)
}
func getByParam(c *gin.Context){
	id := c.Param("id")
	for _,t:= range todos{
		if t.ID == id{
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, id)
}
func main(){
   router := gin.Default()
   router.GET("/todos",getTodos)
   router.POST("/addTodo",addTodos)
   router.POST("/getById",getById)
   router.GET("/getByParam/:id",getByParam)
   router.Run("localhost:9092")
}