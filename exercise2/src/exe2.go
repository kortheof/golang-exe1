package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	Name, Surname, Title string
	Salary               uint64
}

//Constructor function that returns an instance of an Employee struct
func newPerson(s []string) Employee {
	//Convert the salary to uint as the Employee struct requires
	uintSalary, _ := strconv.ParseUint(s[3], 10, 64)
	return Employee{
		Name:    s[0],
		Surname: s[1],
		Title:   s[2],
		Salary:  uintSalary,
	}
}

func main() {
	csvFile, err := os.Open("devops.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	//A slice of Employee struct, every index will hold name, surname, title and salary
	var person []Employee

	//reader := csv.NewReader(bufio.NewReader(csvFile))
	reader := csv.NewReader(csvFile)
	//Read the first line to exclude it
	reader.Read()

	//Parse the input csv file and create a new Employee instance from each line
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		person = append(person, newPerson(line))
	}
	fmt.Println(person[0].Salary)

	//Panw sto slice apo Employee structs, 8a xtisw me8odous na upologizoun ta results.
}
