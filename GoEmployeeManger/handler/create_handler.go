package handler

import (
    "net/http"
    "GoEmployeeManger/model"
    "GoEmployeeManger/service"
    "GoEmployeeManger/config"
    "github.com/gin-gonic/gin"
)

// CreateEmployeeHandler handles POST requests to insert a new employee.
func CreateEmployee(c *gin.Context) {
    // Define a variable to hold the employee data
    var employee model.Employee

    // Bind the JSON data from the request body to the employee struct
    if err := c.ShouldBindJSON(&employee); err != nil {
        // If there is an error in binding, return a Bad Request response
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Call the service function to create the employee in the database
    // Pass the database connection and the employee data
    if err := service.CreateEmployee(config.DB, &employee); err != nil {
        // If there is an error in creating the employee, return an Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // If the employee is successfully created, return a Created response with the employee data
    c.JSON(http.StatusCreated, employee)
}
