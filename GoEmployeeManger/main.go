// main.go

package main

import (
    "GoEmployeeManger/config"
	"GoEmployeeManger/router"
)

func main() {
    config.InitDB()
     r := router.SetupRouter()
     r.Run("localhost:8080")
	
}
