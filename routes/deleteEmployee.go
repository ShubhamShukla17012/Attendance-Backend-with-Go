package routes

import (
	"attandance/models"
	"attandance/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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
