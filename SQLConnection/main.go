package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Starting service")
	// sql.Register("mysql", &MySQLDriver{})
	db,err:=sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s","root","2109@nihira","localhost:3306","training"))
	if err!=nil{
		fmt.Println(fmt.Errorf("Error opening SQL,%s",err.Error()))
		return
	}
	res,err := db.Exec("insert into training.Books values(20,'black',30.5,20.4)")
	if err!=nil{
		fmt.Println(fmt.Errorf("Returned error from SQL EXEC command:%s",err.Error()))
		return
	}
	count,err:=res.RowsAffected()
	if err!=nil{
		fmt.Println(fmt.Errorf("Error at RowsAffected:%s",err.Error()))
		return
	}
	fmt.Println("Successfully added new record into DB:",count)
}