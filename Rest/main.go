package main

import (
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
func main(){
	router:=gin.Default()
	router.GET("/getId",gettingEmployeed)
	router.Run("localhost:9090")
}
func gettingEmployeed(c *gin.Context){
	c.IndentedJSON(http.StatusOK,todos)
}