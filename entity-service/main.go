package main

import (
	"entity-service/controller"
	"entity-service/database"
	"entity-service/repository"
	"entity-service/server"
	"entity-service/service"

	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}
}

func main() {
	d := database.New()

	r := repository.New(d)
	s := service.New(r)
	c := controller.New(s)

	ss := server.New(c)

	err := d.Run()
	if err != nil {
		panic(err)
	}

	err = ss.Run()
	if err != nil {
		panic(err)
	}
}
