package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammadinasab-dev/employee-salary/configuration"
	"github.com/mohammadinasab-dev/employee-salary/data"
)

var handler *EmployeeSalaryHandler

func RunApi(path string) error {

	config, err := configuration.LoadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	db, err := data.CreateDBConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	addr := config.ServerAddress
	mux := mux.NewRouter()
	RunAPIOnRouter(mux, db)
	log.Println("Server Started")
	return http.ListenAndServe(addr, mux)
}

func RunAPIOnRouter(r *mux.Router, db *data.SQLHandler) {
	handler = NewEmployeeSalaryHandler(db)

	r.HandleFunc("/salary", handler.setsalary).Methods("POST")
	r.HandleFunc("/salary", handler.getSalaries).Methods("GET")
	r.HandleFunc("/salary", handler.updatesalary).Methods("PUT")
	r.HandleFunc("/salary", handler.deleteSalary).Methods("DELETE")
	r.HandleFunc("/employee", handler.newEmployee).Methods("POST")
	r.HandleFunc("/employee", handler.getEmploiees).Methods("GET")
	r.HandleFunc("/employee", handler.updateEmployee).Methods("PUT")
	r.HandleFunc("/employee", handler.deleteEmployee).Methods("DELETE")

}
