package routes

import (
	"attandance/pkg"
	"fmt"
	"net/http"
)

func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")
	emp, err := pkg.FindEmployee(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("User not Found :%v", err), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(emp.Jsonify())
}
