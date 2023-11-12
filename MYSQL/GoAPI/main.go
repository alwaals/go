package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Books struct {
	Pages  int32   `json:"pages"`
	Color  string  `json:"color"`
	Height float32 `json:"height"`
	Width  float32 `json:"width"`
}

const tableName = "training.books"

func main() {
	fmt.Println("Strting service")
	dbConn, err := dbConn()
	if err != nil {
		fmt.Println("Error during DB connection:", err.Error())
		return
	}
	defer dbConn.Close()

	dbConn.SetMaxOpenConns(20)
	dbConn.SetMaxIdleConns(20)
	dbConn.SetConnMaxLifetime(time.Minute*10)

	router := gin.Default()
	router.GET("/getbooks", GetBooksHandler(dbConn))
	router.POST("/addbooks", AddBooksHandler(dbConn))
	router.POST("/deletebooks", DeleteBooksHandler(dbConn))
	router.Run("localhost:9090")
}
func DeleteBooksHandler(conn *sql.DB)func(c *gin.Context){
	return func(c *gin.Context){
		ctx,ctxCancel:=context.WithTimeout(context.Background(),time.Second*5)
		defer ctxCancel()
	
		var book Books
		err:=c.BindJSON(&book)
		if err!=nil{
			c.IndentedJSON(http.StatusInternalServerError,err.Error())
			return
		}
		fmt.Println(fmt.Sprintf("%v",book))
		status:=deleteRecord(conn,ctx, book.Pages)
		if !strings.Contains(status, "Success") {
			c.IndentedJSON(http.StatusInternalServerError,status)
			return
		}
		c.IndentedJSON(http.StatusOK,status)
		return
	}
}
func GetBooksHandler(dbConn *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx,ctxCancel:=context.WithTimeout(context.Background(),time.Second*5)
		defer ctxCancel()
	
		// var books []Books
		books, err := fetchRecords(dbConn,ctx)
		if err != "" {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}
		c.IndentedJSON(http.StatusOK, &books)
		return
	}
}
func AddBooksHandler(dbConn *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx,ctxCancel:=context.WithTimeout(context.Background(),time.Second*5)
		defer ctxCancel()
	
		var books Books
		err := c.BindJSON(&books)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}
		errS := insertRecord(dbConn,ctx, &books)
		if !strings.Contains(errS, "Success") {
			c.IndentedJSON(http.StatusBadRequest, errS)
			return
		}
		c.IndentedJSON(http.StatusOK, &books)
		return
	}
}
func dbConn() (*sql.DB, error) {
	//return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "2109@nihira", "localhost:3306", "training"))
	
	return sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s","admin","2109nihira","training.cx5wsl5g4edg.us-east-1.rds.amazonaws.com:3306","training"))
}
func insertRecord(sql *sql.DB,ctx context.Context, book *Books) string {
	fmt.Println(book.Pages,book.Color,book.Height,book.Width)
	res, err := sql.ExecContext(ctx, fmt.Sprintf("INSERT into %s VALUES (%v,\"%s\",%v,%v)", tableName, book.Pages, book.Color, book.Height, book.Width))
	if err != nil {
		return fmt.Sprintf("Error returned from sql Exec %s", err.Error())
	}
	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Sprint("Error returned from RowsEffected %s ", err.Error())
	}
	return fmt.Sprintf("Successfully added new record into database")
}
func fetchRecords(sql *sql.DB,ctx context.Context) ([]Books, string) {
	var books []Books
	fmt.Println("Reached fetchRecords with DbConn:", sql.Stats())
	rows, err := sql.QueryContext(ctx,fmt.Sprintf("Select * from %s",tableName))
	if err != nil {
		return books, fmt.Sprintf("Error returned from sql Exec : %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var b Books
		if err := rows.Scan(&b.Pages, &b.Color, &b.Height, &b.Width); err != nil {
			return books, fmt.Sprintf("Error during rows.Scan:%s", err.Error())
		}
		books = append(books, b)
	}
	return books, ""
}
func deleteRecord(sql *sql.DB, ctx context.Context,pages int32) string {
	fmt.Println("Reached deleteRecord handler!")
	res, err := sql.ExecContext(ctx,fmt.Sprintf("delete from %s where pages = %v", tableName, pages))
	if err != nil {
		return fmt.Sprintf("Error returned from sqlExec:%s", err.Error())
	}
	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Sprintf("Error returned from RowsEffected %s:", err.Error())
	}
	return "Successfully Deleted the record from database"
}
