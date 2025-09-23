package main

import (
	"attandance/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /employee", routes.CreateEmployeeHandler)
	mux.HandleFunc("GET /employee/{id}", routes.GetEmployeeHandler)
	mux.HandleFunc("GET /employees", routes.GetAllEmployeeHandler)
	mux.HandleFunc("DELETE /employee/{id}", routes.DeleteEmployee)
	// mux.HandleFunc("PATCH /employee/login/{id}", loginHandler)
	// mux.HandleFunc("PATCH /epmloyee/logout/{id}", logOutHandlar)
	fmt.Println("Server running on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
