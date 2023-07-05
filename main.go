package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func main() {

	db, err := sql.Open("mysql", "root:vuvietduy1234@tcp(localhost:3306)/todo")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected")
	}

	//an íntance of our todo struct
	item := todo{1, "Golang", false, time.Now(), time.Now()}
	//convert object to json

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(item)
	})

	//app.Post("/items", createItem(db))
	app.Get("/items", getListTodo(db))
	app.Get("/items/:id", readItemById(db))

	log.Fatal(app.Listen(":3000"))
}

//func createItem(db *sql.DB) fiber.Handler {
//	return func(c *fiber.Ctx) error {
//		item := todo{1, "Golang", false, time.Now(), time.Now()}
//
//		return c.JSON(item)
//	}
//}

func readItemById(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dataItem todo
		id := c.Params("id")

		fmt.Println(id)
		//query := "SELECT * FROM todo_item WHERE id = ? "

		row, err := db.Query("SELECT * FROM todo_item WHERE id = ? ", id)
		if err != nil {
			fmt.Println(err)
		}
		row.Next()
		err = row.Scan(&dataItem.Id, &dataItem.Title, &dataItem.Status, &dataItem.CreatedAt, &dataItem.UpdateAt)
		if err != nil {
			fmt.Println(err)
		}
		return c.JSON(dataItem)
	}
}

func getListTodo(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dataItem []todo
		var item todo

		//query := "SELECT * FROM todo_item WHERE id = ? "

		rows, err := db.Query("SELECT * FROM todo_item")
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			rows.Scan(&item.Id, &item.Title, &item.Status, &item.CreatedAt, &item.UpdateAt)
			dataItem = append(dataItem, item)
		}

		return c.JSON(dataItem)
	}
}
