package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type Message struct{
	Id int `json:"id"`
	Msg string `json:"msg"`
}

var db *sql.DB

func init() {
	dbQuery := `
	CREATE TABLE IF NOT EXISTS Messages(
	id INT AUTO_INCREMENT PRIMARY KEY,
	message TEXT
	);`
	if _, err := db.Exec(dbQuery); err != nil {
		log.Fatal("Error Creating Table : ", err)
	}
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.BindEnv("DB_HOST")	
	viper.BindEnv("DB_USER")	
	viper.BindEnv("DB_PASSWORD")	
	viper.BindEnv("DB")	
	if err := viper.ReadInConfig() ; err != nil{
		panic("Error Reading The .env file ! 🔺 Log -> " + err.Error())
	}
}

func main(){
	fmt.Print("Hey There 🧑‍🦳 Its me Message Mate 🧒 Share Me A message  🥇")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
	viper.GetString("DB_HOST"),
	viper.GetString("DB_USER"),
	viper.GetString("DB_PASSWORD"),
	viper.GetString("DB_NAME"),
	)
	var err error 
	db , err = sql.Open("mysql",dsn)
	if err != nil {
		log.Fatal(" 🔺 Error Establishing Database Connection -> ", err)
	}
	defer db.Close()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
        rows, err := db.Query("SELECT id, message FROM messages")
        if err != nil {
            return c.Status(500).SendString("Error querying database")
        }
        defer rows.Close()

        var messages []Message
        for rows.Next() {
            var msg Message
            if err := rows.Scan(&msg.ID, &msg.Message); err != nil {
                return c.Status(500).SendString("Error scanning row")
            }
            messages = append(messages, msg)
        }
        return c.Render("index", fiber.Map{
            "Messages": messages,
        })
    })

    app.Post("/submit", func(c *fiber.Ctx) error {
        newMessage := c.FormValue("new_message")

        _, err := db.Exec("INSERT INTO messages (message) VALUES (?)", newMessage)
        if err != nil {
            return c.Status(500).SendString("Error inserting message")
        }
        return c.JSON(fiber.Map{"message": newMessage})
    })

    // Start the Fiber app
    log.Fatal(app.Listen(fmt.Sprintf(":%s",3000)))
}




















