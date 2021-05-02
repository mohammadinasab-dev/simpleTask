package data

import "time"

type Employee struct {
	EmployeeID int        `gorm:"primary_key" json:"empid"`
	UserName   string     `gorm:"NOT NULL; UNIQUE" json:"username"`
	Job        string     `gorm:"NOT NULL" json:"job"`
	WorkTime   int        `gorm:"NOT NULL" json:"worktime"`
	Salary     int        `gorm:"NOT NULL" json:"salary"`
	CreatedAt  time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW()" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW(); ON UPDATE now()" json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleated_at"`
	//Salary     Salary     `gorm:"ForeignKey:SalaryID"`
}

type Salary struct {
	SalaryID  int        `gorm:"primary_key" json:"salaryid"`
	JobTitle  string     `gorm:"NOT NULL" json:"jobtitle"`
	PriceUnit int        `json:"priceunit"`
	CreatedAt time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW(); ON UPDATE now()" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleated_at"`
}
