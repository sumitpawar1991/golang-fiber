package controller

import (
	"log"
	"my-fiber-app/server/database"
	"my-fiber-app/server/model"

	"github.com/gofiber/fiber/v3"
)

func BlogList(c fiber.Ctx) error {

	context := fiber.Map{
		"status":  "ok",
		"message": "Blog List",
	}
	// c.Status(200)
	// return c.JSON(context)

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog records"] = records
	c.Status(200)
	return c.JSON(context)

}

func BlogCreate(c fiber.Ctx) error {

	context := fiber.Map{
		"status":  "ok",
		"message": "Blog List",
	}

	var record model.Blog

	if err := c.Bind().Body(&record); err != nil {
		log.Println("error in parsing the request")
		context["status"] = "false"
		context["message"] = "something went wrong"
	}

	result := database.DBConn.Create(&record)

	if result.Error != nil {
		log.Println("Error in saving data")
		context["status"] = "false"
		context["message"] = "something went wrong. coudnt insert data in the table"
	}

	context["message"] = "Data inserted successfully"
	context["data"] = record
	c.Status(201)
	return c.JSON(context)

}

func BlogUpdate(c fiber.Ctx) error {

	context := fiber.Map{
		"status":  "ok",
		"message": "Blog Update",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.Bind().Body(&record); err != nil {
		log.Println("Error in parsing data")
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data")
	}

	context["data"] = record
	context["message"] = "Record Updated Successfully"

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c fiber.Ctx) error {

	context := fiber.Map{
		"status":  "ok",
		"message": "Blog delete",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["message"] = "Record not found"

		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		context["message"] = "Someting went wrong"
		c.Status(400)
	}

	context["message"] = "Record Deleted Successfully"
	c.Status(200)
	return c.JSON(context)
}
