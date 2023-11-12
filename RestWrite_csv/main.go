package main

import (
	//"bytes"
	//"encoding/csv"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){

	// e:=Employee{empId: 101,empName: "Swetha",empSal: 20000,empExp: 5}
	// es,err:=os.Create("emp.csv")
	// if err!=nil{
	// 	fmt.Println(fmt.Errorf("Unable to create the CSV file%s",err.Error()))
	// 	return
	// }
	// defer es.Close()
	// _,errL:=es.WriteString(fmt.Sprintf("EmployeeID  EmployeeName\n"))
	// _,errL=es.WriteString(fmt.Sprintf("%v \t\t%s",e.empId,e.empName))
	// if errL!=nil{
	// 	fmt.Println(fmt.Errorf("Unable to write into csv %s",errL.Error()))
	// 	return
	// }
	// fmt.Println("Done writing into file!")

	//
	readFile,errL:=os.Open("./read.csv")
	if errL!=nil{
		fmt.Println(fmt.Errorf("Unable to read the csv %s",errL.Error()))
		return
	}
	defer readFile.Close()

	rd:=bufio.NewReader(readFile)
	for{
		line,err:=rd.ReadString('\n')
		if err!=nil{
			if err==io.EOF{
				break
			}
			fmt.Println(fmt.Errorf("Unable to read the csv %s",errL.Error()))
		    return
		}
		fmt.Println(line)
	}
	
}

type Employee struct{
	empId int `json:'EmpID'`
	empName string `json:'EmpName'`
	empSal int `json:'EmpSalary'`
	empExp int `json:'EmpExperiance'`
}