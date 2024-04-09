package main

import (
	"fmt"
	middleware "ginjing/Middleware"
	"ginjing/controllers/Todo"
	"ginjing/db"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Some error occured. Err: %s", err)
		return
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "migrate" {
			fmt.Println("Migrating database to turso..")
			db.Migrate()
			fmt.Println("Database migrated 👍.")
		} else {
			fmt.Println(
				strings.TrimSpace(
					`Unknown Command!
---------------
Accepted commands:
- migrate				: Migrates the database to turso`,
				),
			)
		}
	} else {
		app := iris.New()

		app.UseRouter(recover.New())

		todoAPI := app.Party("/todo")
		{
			todoAPI.Use(iris.Compression)

			todoAPI.Get("/", Todo.List)
			todoAPI.Post("/", Todo.Create)
		}

		todoProtectedAPI := app.Party("/protected/todo")
		{
			todoProtectedAPI.Use(iris.Compression)
			todoProtectedAPI.Use(middleware.AuthMiddleware)

			todoProtectedAPI.Get("/", Todo.List)
		}

		app.Listen(":5000")
	}
}
