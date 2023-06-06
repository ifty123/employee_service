package main

import (
	"flag"
	"latihan_service/database"
	"latihan_service/database/migration"
	"latihan_service/database/seeder"
	"latihan_service/internal/factory"
	"latihan_service/internal/http"
	"latihan_service/internal/middleware"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database.GetConnection()
}

func main() {
	database.CreateConnection()

	//ini untuk migrate dan seeder
	var m, s string

	//membaca dari command line
	flag.StringVar(
		&m,
		"m",
		"none",
		`this argument for check if user want to migrate table, rollback table, or status migration
to use this flag:
	use -m=migrate for migrate table
	use -m=rollback for rollback table
	use -m=status for get status migration`,
	)

	flag.StringVar(
		&s,
		"s",
		"none",
		`this argument for check if user want to seed table
to use this flag:
	use -s=all to seed all table`,
	)
	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
	} else if m == "rollback" {
		migration.Rollback()
	} else if m == "status" {
		migration.Status()
	}

	if s == "all" {
		seeder.NewSeeder().DeleteAll()
		seeder.NewSeeder().SeedAll()
	}

	//factory database
	f := factory.NewFactory()
	e := echo.New()
	mid := middleware.NewMidleware()

	e.Use(mid.CORS)
	mid.LogMiddlewares(e)

	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
