package routes

import (
	"attandance/models"
	"attandance/pkg"
	"attandance/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// create employee
func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR", "failed to read raw body", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.NewAPIError(http.StatusBadRequest, err.Error()).Jsonify())
		return
	}
	defer r.Body.Close()

	var emp models.Employee
	err = json.Unmarshal(rawBody, &emp)
	if err != nil {
		log.Println("ERROR", "failed to unmarshall body", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.NewAPIError(http.StatusBadRequest, err.Error()).Jsonify())
		return
	}
	// TODO : Set Login time and logout time zero

	// Add a valid UUID to employee
	emp.ID = uuid.NewString()

	if err := emp.Validate(false); err != nil {
		log.Println("ERROR", "failed to read body", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.NewAPIError(http.StatusBadRequest, err.Error()).Jsonify())
		return
	}

	err = pkg.AddEmployee(&emp)
	if err != nil {
		log.Println("ERROR", "failed to Add employee to file", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.NewAPIError(http.StatusInternalServerError, err.Error()).Jsonify())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(emp.Jsonify())
}

// send all employees to User
func GetAllEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(utils.EmployeeFile)
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

// get employee by id
func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")
	emp, err := pkg.FindEmployee(ID)
	fmt.Print(err)
	w.WriteHeader(http.StatusOK)
	w.Write(emp.Jsonify())
}

// TODO:: write Delete Employee By ID
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	empID := r.PathValue("id")
	employeesData, err := os.ReadFile(utils.EmployeeFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}
	var employees []models.Employee

	if err := json.Unmarshal(employeesData, &employees); err != nil {
		http.Error(w, "Error parsing employee data", http.StatusInternalServerError)
		return
	}
	updatedEmployees := []models.Employee{}
	found := false
	for _, emp := range employees {
		if emp.ID != empID {
			updatedEmployees = append(updatedEmployees, emp)
		} else {
			found = true
		}
	}
	if !found {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	updatedData, err := json.MarshalIndent(updatedEmployees, "", "  ")
	if err != nil {
		http.Error(w, "Error writing updated data", http.StatusInternalServerError)
		return
	}
	if err := os.WriteFile(utils.EmployeeFile, updatedData, 0644); err != nil {
		http.Error(w, "Error saving updated file", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
