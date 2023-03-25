package main

import (
	"Belajar-Golang-7/db"
	"Belajar-Golang-7/router"
)

func main() {
	db.StartDB()
	router.StartServer().Run(":8080")
}
