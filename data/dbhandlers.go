package data

import "errors"

func (handler *SQLHandler) DBInsertSalary(salary Salary) error {
	result := handler.DB.Debug().Create(&salary)
	if result.Error != nil {
		return errors.New("failed to add new salary")
	}
	if result.RowsAffected == 0 {
		return errors.New("no row effected")
	}
	return nil
}

func (handler *SQLHandler) DBGetSalariess() ([]Salary, error) {

	salariess := []Salary{}
	result := handler.DB.Debug().Find(&salariess)
	if result.Error != nil {
		return nil, errors.New("failed to get salariess")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no row effected")
	}
	return salariess, nil
}

func (handler *SQLHandler) DBUpdateSalary(salary Salary) error {
	result := handler.DB.Debug().Model(&salary).Where("job_title = ?", salary.JobTitle).Update(&salary)
	if result.Error != nil {
		return errors.New("failed to update salary")
	}
	if result.RowsAffected == 0 {
		return errors.New("no row effected")
	}
	return nil
}

func (handler *SQLHandler) DBDeleteSalary(salary Salary) error {

	result := handler.DB.Debug().Where("job_title = ?", salary.JobTitle).Delete(&Salary{})
	if result.Error != nil {
		return errors.New("failed to delete salary")
	}
	if result.RowsAffected == 0 {
		return errors.New("no row effected")
	}
	return nil
}

func (handler *SQLHandler) DBFindSalary(jobTitle string) (int, error) {
	var salary Salary
	result := handler.DB.Debug().Where("job_title = ?", jobTitle).First(&salary)
	if result.Error != nil {
		return 0, errors.New("failed to get salary.priceunit")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("no row effected")
	}
	return salary.PriceUnit, nil
}

func (handler *SQLHandler) DBInsertEmployee(employee Employee) (Employee, error) {
	priceUnit, err := handler.DBFindSalary(employee.Job)
	if err != nil {
		return Employee{}, err
	}
	employee.Salary = employee.WorkTime * priceUnit
	result := handler.DB.Debug().Create(&employee)
	if result.Error != nil {
		return Employee{}, errors.New("failed to add new employee")
	}
	if result.RowsAffected == 0 {
		return Employee{}, errors.New("no row effected")
	}
	return employee, nil
}

func (handler *SQLHandler) DBGetEmploiees() ([]Employee, error) {

	emploiees := []Employee{}
	result := handler.DB.Debug().Find(&emploiees)
	if result.Error != nil {
		return nil, errors.New("failed to get salariess")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no row effected")
	}
	return emploiees, nil
}

func (handler *SQLHandler) DBUpdateEmployee(employee Employee) error {
	result := handler.DB.Debug().Model(&employee).Where("user_name = ?", employee.UserName).Update(&employee)
	if result.Error != nil {
		return errors.New("failed to update employee")
	}
	if result.RowsAffected == 0 {
		return errors.New("no row effected")
	}
	return nil
}

func (handler *SQLHandler) DBDeleteEmployee(employee Employee) error {

	result := handler.DB.Debug().Where("user_name = ?", employee.UserName).Delete(&Employee{})
	if result.Error != nil {
		return errors.New("failed to delete employee")
	}
	if result.RowsAffected == 0 {
		return errors.New("no row effected")
	}
	return nil
}
