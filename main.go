package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Car struct {
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Color  string `json:"color"`
	Number string `json:"number"`
}

var cars []*Car

func createCarHandler(c *fiber.Ctx) error {
	var car Car
	err := c.BodyParser(&car)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	cars = append(cars, &car)
	return c.SendStatus(fiber.StatusCreated)
}

func getCarHandler(c *fiber.Ctx) error {
	number := c.Query("number")
	for _, car := range cars {
		if car.Number == number {
			return c.JSON(car)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}

func updateCarHandler(c *fiber.Ctx) error {
	number := c.Query("number")
	var newCar Car
	err := c.BodyParser(&newCar)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for _, car := range cars {
		if car.Number == number {
			car.Brand = newCar.Brand
			car.Model = newCar.Model
			car.Color = newCar.Color
			car.Number = newCar.Number
			return c.SendStatus(fiber.StatusOK)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}

func deleteCarHandler(c *fiber.Ctx) error {
	number := c.Query("number")
	for i, car := range cars {
		if car.Number == number {
			cars = append(cars[:i], cars[i+1:]...)
			return c.SendStatus(fiber.StatusOK)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Not Found")
}

func main() {
	app := fiber.New()

	app.Post("/car", createCarHandler)
	app.Get("/car", getCarHandler)
	app.Put("/car", updateCarHandler)
	app.Delete("/car", deleteCarHandler)

	fmt.Println("Server is running on port 8080...")
	app.Listen(":8080")
}
