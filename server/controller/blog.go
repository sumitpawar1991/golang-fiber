package controller

import (
	"log"
	"my-fiber-app/server/database"
	"my-fiber-app/server/model"

	import "github.com/go-playground/validator/v10"
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

	context["blog_records"] = records
	c.Status(200)
	return c.JSON(context)

}

func BlogDetail(c fiber.Ctx) error {

	context := fiber.Map{
		"status":  "",
		"message": "",
	}
	// c.Status(200)
	// return c.JSON(context)

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["message"] = "Record not found"

		c.Status(404)
		return c.JSON(context)
	}

	context["record"] = record
	context["status"] = "ok"
	context["message"] = "Record fetch successfully"
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

// func BlogCreate(c fiber.Ctx) error {
// 	var req BlogRequest

// 	if err := c.Bind().Body(&req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "Invalid request body",
// 		})
// 	}

// 	if err := validate.Struct(req); err != nil {

// 		errors := make(map[string]string)

// 		for _, e := range err.(validator.ValidationErrors) {
// 			switch e.Field() {
// 			case "Title":
// 				errors["title"] = "Title is required"
// 			case "Post":
// 				errors["post"] = "Post must be at least 10 characters"
// 			}
// 		}

// 		return c.Status(400).JSON(fiber.Map{
// 			"status": "error",
// 			"errors": errors,
// 		})
// 	}

// 	// Insert into database
// 	_, err := DB.Exec(
// 		"INSERT INTO blogs(title, post) VALUES (?, ?)",
// 		req.Title,
// 		req.Post,
// 	)

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status":  "ok",
// 		"message": "Blog created successfully",
// 	})

// }
