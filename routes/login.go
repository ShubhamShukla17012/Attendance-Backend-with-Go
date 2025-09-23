package routes

import (
	"attandance/models"
	"attandance/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	empID := r.PathValue("id")

	// Read employee data from file
	employeeData, err := utils.FileReadOpration()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	var employees []models.Employee
	err = json.Unmarshal(employeeData, &employees)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error unmarshalling data: %v", err), http.StatusInternalServerError)
		return
	}

	// Flag to check if employee was found
	found := false

	// Update login time
	for i := range employees {
		if employees[i].ID == empID {
			employees[i].LogInTime = time.Now()
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// Marshal updated data
	updatedData, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshalling updated data: %v", err), http.StatusInternalServerError)
		return
	}

	// Write back to file
	err = os.WriteFile(utils.EmployeeFile, updatedData, 0666)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing to file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login time updated successfully"))
}
