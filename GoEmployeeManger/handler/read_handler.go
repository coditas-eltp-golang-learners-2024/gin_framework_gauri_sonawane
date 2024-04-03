package handler

import (
	"GoEmployeeManger/config"
	"GoEmployeeManger/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetEmployeesHandler handles GET requests to retrieve all employees.
func GetEmployees(c *gin.Context) {
    employees, err := service.GetEmployees(config.DB) // Pass db connection to service function
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, employees)
}
