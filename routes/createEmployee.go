package routes

import (
	"attandance/models"
	"attandance/pkg"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

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
