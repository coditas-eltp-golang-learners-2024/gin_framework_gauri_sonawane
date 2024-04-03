// router/router.go

package router

import (
    "github.com/gin-gonic/gin"
    "GoEmployeeManger/handler"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Route to get all employees
    r.GET("/employees", handler.GetEmployees)

    // Route to create a new employee
    r.POST("/employees", handler.CreateEmployee)

    return r
}

