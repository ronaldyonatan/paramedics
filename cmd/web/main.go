package main

import (
	"log"

	"github.com/fernandojec/assignment-2/domain/users"
	"github.com/fernandojec/assignment-2/pkg/dbconnect"
	"github.com/fernandojec/assignment-2/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	// err = godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		// fmt.Printf("Error load env:%v", err)
		log.Fatalf("Cannot get env :%v", err)

	}
	dbx, err := dbconnect.ConnectSqlx(dbconnect.DBConfig{
		Host:       utils.GetEnv("POSTGRES_HOST"),
		Port:       utils.GetEnv("POSTGRES_PORT"),
		Dbname:     utils.GetEnv("POSTGRES_DBNAME"),
		Dbuser:     utils.GetEnv("POSTGRES_USER"),
		Dbpassword: utils.GetEnv("POSTGRES_PASSWORD"),
		Sslmode:    utils.GetEnv("POSTGRES_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("Cannot connect to DB:%v", err)
	}

	engine := html.New("../../views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/web",
		},
		StatusCode: 301,
	}))
	web := app.Group("web")
	users.RouterWebInit(web, dbx)

	log.Fatal(app.Listen(":3001"))
}
