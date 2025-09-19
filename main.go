package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Employee struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Grade      string    `json:"grade"`
	LogInTime  time.Time `json:"log_out_time,omitempty"`
	LogOutTime time.Time `json:"log_in_time,omitempty"`
}

var employees []Employee

func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if emp.ID != "" {
		emp.ID = ""
	}
	if emp.Grade == "" || len(emp.Grade) > 2 || emp.Name == "" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Please write mandatory feild")
		return
	}
	employees = append(employees, emp)

	file, err := os.Create("employees.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(employees)
	if err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

	fmt.Println("Employee saved to employees.json")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)

}
func getAllEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("employees.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	if !json.Valid(data) {
		http.Error(w, "Invalid JSON in file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /employee", createEmployeeHandler)
	// mux.HandleFunc("/GET/employee/{id}", getEmployeeHandler)
	mux.HandleFunc("GET /Employees", getAllEmployeeHandler)
	// mux.HandleFunc("/DELETE/employee/{id}", deleteEmployee)
	// mux.HandleFunc("/PATCH/employee/login/{id}", loginHandler)
	// mux.HandleFunc("/PATCH/epmloyee/logout/{id}", logOutHandlar)
	fmt.Println("Server running on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
