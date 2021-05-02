package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mohammadinasab-dev/employee-salary/data"
	"github.com/mohammadinasab-dev/employee-salary/response"
)

type EmployeeSalaryHandler struct {
	dbhandler *data.SQLHandler
}

func NewEmployeeSalaryHandler(db *data.SQLHandler) *EmployeeSalaryHandler {
	return &EmployeeSalaryHandler{
		dbhandler: db,
	}
}

func (handler EmployeeSalaryHandler) setsalary(w http.ResponseWriter, r *http.Request) {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.ERROR(w, "false", "Error to read request body", http.StatusInternalServerError, err)
	}
	salary := data.Salary{}
	err = json.Unmarshal(jsonbyte, &salary)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusBadRequest, err)
	}

	err = handler.dbhandler.DBInsertSalary(salary)
	if err != nil {
		response.ERROR(w, "false", "Error of database create query", http.StatusInternalServerError, err)
	}
	err = response.JSON(w, "true", "new salary set successfully", http.StatusOK, "")
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusInternalServerError, err)
	}
}

func (handler EmployeeSalaryHandler) getSalaries(w http.ResponseWriter, r *http.Request) {
	salariess, err := handler.dbhandler.DBGetSalariess()
	if err != nil {
		response.ERROR(w, "false", "Error of database find query", http.StatusInternalServerError, err)
	}
	err = response.JSON(w, "true", "books load successfully", http.StatusOK, salariess)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusInternalServerError, err)
	}
}

func (handler EmployeeSalaryHandler) updatesalary(w http.ResponseWriter, r *http.Request) {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.ERROR(w, "false", "Error to read request body", http.StatusInternalServerError, err)
	}
	salary := data.Salary{}
	err = json.Unmarshal(jsonbyte, &salary)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusBadRequest, err)
	}
	err = handler.dbhandler.DBUpdateSalary(salary)
	if err != nil {
		response.ERROR(w, "false", "Error of database no row affected", http.StatusInternalServerError, err)
	}
	err = response.JSON(w, "true", "salary update successfully", http.StatusOK, "")
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusInternalServerError, err)
	}
}

func (handler EmployeeSalaryHandler) deleteSalary(w http.ResponseWriter, r *http.Request) {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.ERROR(w, "false", "Error to read request body", http.StatusInternalServerError, err)
	}
	salary := data.Salary{}
	err = json.Unmarshal(jsonbyte, &salary)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusBadRequest, err)
	}

	err = handler.dbhandler.DBDeleteSalary(salary)
	if err != nil {
		response.ERROR(w, "false", "Error of database (delete) no row affected:", http.StatusInternalServerError, err)
	}
	err = response.JSON(w, "true", "salary deleted successfully", http.StatusOK, "")
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusInternalServerError, err)
	}
}

func (handler EmployeeSalaryHandler) newEmployee(w http.ResponseWriter, r *http.Request) {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.ERROR(w, "false", "Error to read request body", http.StatusInternalServerError, err)
	}
	employee := data.Employee{}
	err = json.Unmarshal(jsonbyte, &employee)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusBadRequest, err)
	}

	employee, err = handler.dbhandler.DBInsertEmployee(employee)
	if err != nil {
		response.ERROR(w, "false", "Error of database create query", http.StatusInternalServerError, err)
	}
	err = response.JSON(w, "true", "new employee add successfully", http.StatusOK, employee)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusInternalServerError, err)
	}
}

func (handler EmployeeSalaryHandler) getEmploiees(w http.ResponseWriter, r *http.Request) {
	emploiees, err := handler.dbhandler.DBGetEmploiees()
	if err != nil {
		response.ERROR(w, "false", "Error of database find query", http.StatusInternalServerError, err)
	}
	err = response.JSON(w, "true", "books load successfully", http.StatusOK, emploiees)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusInternalServerError, err)
	}
}

func (handler EmployeeSalaryHandler) updateEmployee(w http.ResponseWriter, r *http.Request) {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.ERROR(w, "false", "Error to read request body", http.StatusInternalServerError, err)
	}
	employee := data.Employee{}
	err = json.Unmarshal(jsonbyte, &employee)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusBadRequest, err)
	}
	err = handler.dbhandler.DBUpdateEmployee(employee)
	if err != nil {
		response.ERROR(w, "false", "Error of database no row affected", http.StatusInternalServerError, err)
	}
	err = response.JSON(w, "true", "employee update successfully", http.StatusOK, "")
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusInternalServerError, err)
	}
}

func (handler EmployeeSalaryHandler) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.ERROR(w, "false", "Error to read request body", http.StatusInternalServerError, err)
	}
	employee := data.Employee{}
	err = json.Unmarshal(jsonbyte, &employee)
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusBadRequest, err)
	}

	err = handler.dbhandler.DBDeleteEmployee(employee)
	if err != nil {
		response.ERROR(w, "false", "Error of database (delete) no row affected:", http.StatusInternalServerError, err)
	}
	err = response.JSON(w, "true", "employee deleted successfully", http.StatusOK, "")
	if err != nil {
		response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusInternalServerError, err)
	}
}
