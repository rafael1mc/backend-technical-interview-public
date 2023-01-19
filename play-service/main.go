package main

import (
	"play-service/athletes"
	"play-service/controller"
	"play-service/database"
	"play-service/repository"
	"play-service/server"
	"play-service/service"

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

	a := athletes.New()

	r := repository.New(d)
	s := service.New(r, a)
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
