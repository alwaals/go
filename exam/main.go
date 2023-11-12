package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	EmpId   int    `json:"EmployeeID"`
	EmpName string `json:"EmployeeName"`
	EmpSal  int    `json:"EmployeeSalary"`
}

var emp = []Employee{
	{EmpId: 101, EmpName: "Swetha", EmpSal: 2000},
	{EmpId: 102, EmpName: "Swetha123", EmpSal: 3000},
	{EmpId: 103, EmpName: "Swetha456", EmpSal: 5500},
}

func main() {
	router := gin.Default()
	router.GET("/getDetails", employee)
	router.POST("addEmployee",addEmployee)
	router.Run("localhost:8088")
}

func employee(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, emp)
}
func addEmployee(c *gin.Context) {
	var e Employee
	err:=c.BindJSON(&e)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,fmt.Errorf("Error:%s",err.Error()))
		return
	}
	emp = append(emp, e)
	c.IndentedJSON(http.StatusOK, emp)
}
