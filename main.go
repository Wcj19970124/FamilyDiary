package main

import (
	"./models"
	_ "./routers"
)

func main() {
	// beego.Run()
	models.GetUser()

}
