package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Employee struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Title   string `json:"title"`
	Salary  uint64 `json:"salary"`
}

type EmployeeSlice []Employee

//Constructor function that returns an instance of an Employee struct
func newPerson(s []string) (Employee, error) {
	if len(s) >= 5 {
		return Employee{}, fmt.Errorf("More than 4 fields in line %s", s)
	}
	//Convert the salary to uint as the Employee struct requires
	uintSalary, _ := strconv.ParseUint(s[3], 10, 64)
	return Employee{
		Name:    s[0],
		Surname: s[1],
		Title:   s[2],
		Salary:  uintSalary,
	}, nil
}

//Method to calculate the average salary of the input Employees slice
//It has a slice receiver, meaning that every invokation is very cheap on memory
func (e EmployeeSlice) AverageSalary() uint64 {
	total := uint64(0)
	for _, v := range e {
		total += v.Salary
	}
	return total / uint64(len(e))
}

//Method to calculate the maximum salary of the input Employees slice
func (e EmployeeSlice) MaxSalary() uint64 {
	max := uint64(0)
	for _, v := range e {
		if v.Salary > max {
			max = v.Salary
		}
	}
	return max
}

//Method that retrieves the Employees that receive the maximum salary
func (e EmployeeSlice) BiggestSalary() EmployeeSlice {
	maxSal := e.MaxSalary()
	var s EmployeeSlice
	for i, v := range e {
		if v.Salary == maxSal {
			s = append(s, e[i])
		}
	}
	return s
}

//Method to find the number of employees per position
func (e EmployeeSlice) TitleEmployees() map[string]int {
	empTitle := make(map[string]int)
	for _, v := range e {
		_, exists := empTitle[v.Title]
		if exists == false {
			empTitle[v.Title] = 1
		} else {
			empTitle[v.Title] += 1
		}
	}
	return empTitle
}

func main() {
	//Load the defined input csv file
	csvFile, err := os.Open("devops.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	var person EmployeeSlice

	reader := csv.NewReader(csvFile)
	//Read the first line to exclude it
	reader.Read()

	//Parse the file and create a new Employee instance from each line
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		newEmp, err := newPerson(line)
		if err != nil {
			log.Println(err)
			continue
		}
		person = append(person, newEmp)
	}

	//Declare structs to host all the requested statistics
	type AvSal struct {
		AvSal uint64 `json:"average_salary"`
	}

	type BigSal struct {
		BigSal EmployeeSlice `json:"biggest_salary"`
	}

	type Statistics struct {
		AverageSalary AvSal
		EmpPerJob     map[string]int
		BiggestSalary BigSal
	}

	//Calculate statistics based on the created Employees slice and gather the results in a struct
	empStats := Statistics{
		AverageSalary: AvSal{
			AvSal: person.AverageSalary(),
		},
		EmpPerJob: person.TitleEmployees(),
		BiggestSalary: BigSal{
			BigSal: person.BiggestSalary(),
		},
	}

	//Initiate configuration for Web Server to expose the calculated values
	//Define new *ServeMux
	mux := http.NewServeMux()

	//Define new Server
	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	//Define all handler functions
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleNotFound(w, r)
	})

	mux.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		handlePost(w, r, person)
	})

	mux.HandleFunc("/average", func(w http.ResponseWriter, r *http.Request) {
		handleGet(w, r, empStats.AverageSalary)
		//Use the below function call, in order to check how the Web Server responds when json cannot Marshal the input value
		//value := make(chan int)
		//handleGet(w, r, value)
	})

	mux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		handleGet(w, r, empStats.EmpPerJob)
	})

	mux.HandleFunc("/big", func(w http.ResponseWriter, r *http.Request) {
		handleGet(w, r, empStats.BiggestSalary)
	})

	//Start the Web Server
	log.Printf("Web Server started successfully, listening on port %s", server.Addr)
	log.Fatalln("Web Server startup failed with error:", server.ListenAndServe())

}

//Function to handle Web Server generic requests on non-configured endpoints
func handleNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 page not found", http.StatusNotFound)
	//Log an Apache format response - GET /url HTTP1/1 404
	log.Printf("%v %v %v %v", r.Method, r.URL, r.Proto, http.StatusNotFound)
}

//Function to handle Web Server requests on configured endpoints
func handleGet(w http.ResponseWriter, r *http.Request, i interface{}) {
	objJson, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Printf("%v %v %v %v Error: %v", r.Method, r.URL, r.Proto, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Fprintf(w, "%s", string(objJson))
	log.Printf("%v %v %v %v", r.Method, r.URL, r.Proto, http.StatusOK)
}

//Function to handle the json input requests
func handlePost(w http.ResponseWriter, r *http.Request, person EmployeeSlice) {
	var user Employee
	err := json.NewDecoder(r.Body).Decode(&user)
	//Empty input string
	if err == io.EOF {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		log.Printf("%v %v %v %v Error: %v", r.Method, r.URL, r.Proto, http.StatusBadRequest, "Empty input string")
		return
	}
	//Non-decodable input string
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		log.Printf("%v %v %v %v Error: %v", r.Method, r.URL, r.Proto, http.StatusBadRequest, err.Error())
		return
	}
	//Input string was successfully decoded, searching for matches
	for _, value := range person {
		if value.Surname == user.Surname {
			//If a match is found, marshal it in json and return the result to the user
			handleGet(w, r, value)
			return
		}
	}
	//Generic response, no error returned, and no match found
	fmt.Fprintf(w, "[]")
	log.Printf("%v %v %v %v %v", r.Method, r.URL, r.Proto, http.StatusOK, "Match not found")
}
