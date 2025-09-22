package routes

import (
	"attandance/models"
	"attandance/pkg"
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
func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")
	fmt.Println(ID)
	// data, err := os.ReadFile("employees.json")
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	// var emp1  models.Employee
	// json.Unmarshal(data,&emp1)
	// for {
	// 	 {

	// 	}
	// }

}
