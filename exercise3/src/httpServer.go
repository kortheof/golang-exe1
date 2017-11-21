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

/*
var mux map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fanis_result, %q", html.EscapeString(r.URL.Path))
}

//Orizw ta actions tou myHandler --Ti 8a kanei analoga ta calls pou dexetai
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if v, ok := mux[r.URL.String()]; ok {
		v(w, r)
	} else {
		http.NotFound(w, r)
	}
}
*/

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

//Function to convert and print input to json format
func JsonPrint(i interface{}) string {
	objJson, err := json.Marshal(i)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(objJson) + "\n")
	return string(objJson)
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
			fmt.Println(err)
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

	/*
		server := http.Server{
			Addr:    ":8000",
			Handler: &myHandler{},
		}
	*/

	/*
		mux = make(map[string]func(http.ResponseWriter, *http.Request))
		mux["/"] = echoString
		mux["/average"] = echoString
		mux["/employees"] = echoString
		mux["/big"] = echoString
		//mux["/employee"]
	*/

	http.HandleFunc("/average", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", JsonPrint(empStats.AverageSalary))
	})

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", JsonPrint(empStats.EmpPerJob))
	})

	http.HandleFunc("/big", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", JsonPrint(empStats.BiggestSalary))
	})

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		var user Employee
		json.NewDecoder(r.Body).Decode(&user)
		for _, value := range person {
			if value.Surname == user.Surname {
				fmt.Fprintf(w, "%s", JsonPrint(value))

			}
		}
	})

	//server.ListenAndServe()
	http.ListenAndServe(":8000", nil)
}
