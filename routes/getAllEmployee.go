package routes

import (
	"attandance/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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
