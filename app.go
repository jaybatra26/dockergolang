package main

import (
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx)error{
	return c.SendString("this is jaybatra")
}

func setuproutes(app *fiber.App){
	app.Get("/",hello)
}


 

func main(){
	app := fiber.New()

	setuproutes(app)
	
	app.Listen(":6000")

}