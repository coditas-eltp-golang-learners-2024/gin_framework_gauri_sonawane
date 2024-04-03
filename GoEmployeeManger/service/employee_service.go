package service

import (
    "GoEmployeeManger/model"
    "database/sql"
)

// CreateEmployee creates a new employee record in the database.
func CreateEmployee(db *sql.DB, employee *model.Employee) error {
    // Prepare the SQL statement for inserting an employee
    query := "INSERT INTO employees (name, age, salary) VALUES (?, ?, ?)"
    
    // Execute the SQL statement with the employee data
    _, err := db.Exec(query, employee.Name, employee.Age, employee.Salary)
    if err != nil {
        return err
    }

    // If insertion is successful, return nil
    return nil
}

// GetAllEmployees retrieves all employees from the database.
func GetEmployees(db *sql.DB) ([]model.Employee, error) {
    // Query to select all employees from the database
    query := "SELECT * FROM employees"

    // Execute the query
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Initialize a slice to hold the retrieved employees
    var employees []model.Employee

    // Iterate over the rows and scan each row into an Employee struct
    for rows.Next() {
        var employee model.Employee
        err := rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Salary)
        if err != nil {
            return nil, err
        }
        // Append the employee to the slice
        employees = append(employees, employee)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    // Return the retrieved employees
    return employees, nil
}
