package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	// err = godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		// fmt.Printf("Error load env:%v", err)
		log.Fatalf("Cannot get env :%v", err)

	}
	// dbx, err := dbconnect.ConnectSqlx(dbconnect.DBConfig{
	// 	Host:       utils.GetEnv("POSTGRES_HOST"),
	// 	Port:       utils.GetEnv("POSTGRES_PORT"),
	// 	Dbname:     utils.GetEnv("POSTGRES_DBNAME"),
	// 	Dbuser:     utils.GetEnv("POSTGRES_USER"),
	// 	Dbpassword: utils.GetEnv("POSTGRES_PASSWORD"),
	// 	Sslmode:    utils.GetEnv("POSTGRES_SSLMODE"),
	// })
	// if err != nil {
	// 	log.Fatalf("Cannot connect to DB:%v", err)
	// }
	// // _ = dbx
	// app := fiber.New(
	// 	fiber.Config{
	// 		ErrorHandler: customvalidator.HttpErrorHandler,
	// 	},
	// )

	//v1 := app.Group("v1")

	// paramedics.RouterInit(v1, dbx)
	// schedules.RouterInit(v1, dbx)

	//app.Listen(utils.GetEnv("BASE_PORT"))
}
