package postgresdb

import (
	"database/sql"
	

	"github.com/ellipizle/golang-mvc/model"
	"github.com/ellipizle/golang-mvc/pkg/id"
	"github.com/ellipizle/golang-mvc/repository"
	_ "github.com/lib/pq"
	
	// "gopkg.in/mgo.v2/bson"
)

type Postgresdb struct {
	db *sql.DB
}

func New(db *sql.DB) repository.Repository {
	return &Postgresdb{db}
}

func (pg *Postgresdb) AddEmployee(employee *model.Employee) (*model.Employee, error) {
	employee.Id = id.GenerateNewUniqueCode()
	
	datas := "INSERT INTO employees (empid, name, email, salary) VALUES($1, $2, $3, $4,)"
	_, err := pg.db.Query(datas, employee.Id, employee.Name, employee.Email, employee.Salary)
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (pg *Postgresdb) EditEmployee(employee *model.Employee) (*model.Employee, error) {
		data := "UPDATE employees SET name=$2, email=$4, salary=$5 WHERE empid=$1"
	_, err := pg.db.Query(data, employee.Id, employee.Name, employee.Email, employee.Salary)
	if err != nil {
		return employee, err
	}
	return employee, nil

}
func (pg *Postgresdb) GetEmployee(id string) (*model.Employee, error) {
	employee := new(model.Employee)
	data := "SELECT * FROM employees WHERE empid=$1"
	// data := "SELECT * FROM employees"
	rows, err := pg.db.Query(data, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Salary)
		if err != nil {
			panic(err)
		}
	}
	return employee, nil
}
func (pg *Postgresdb) DeleteEmployee(id string) error {
	data := "DELETE FROM employees WHERE empid=$1"
	_, err := pg.db.Query(data, id)
	if err != nil {
		return err
	}
	return nil

}

func (pg *Postgresdb) GetAllEmployee() ([]*model.Employee, error) {
	var employee []*model.Employee
	data := "SELECT * FROM employees"
	rows, err := pg.db.Query(data)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		newEmployee := model.Employee{}
		err = rows.Scan(&newEmployee.Id, &newEmployee.Name, &newEmployee.Email, &newEmployee.Salary)
		if err != nil {
			panic(err)
		}
		employee = append(employee, &newEmployee)
	}
	return employee, nil

}
