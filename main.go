package main

import (
	"fmt"

	"example.com/seen-tech-rtx/DBManager"
	"example.com/seen-tech-rtx/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func SetupRoutes(app *fiber.App) {
	Routes.BranchRoute(app.Group("/dakbox"))
}

func main() {
	fmt.Println("Hello DakBox")

	fmt.Print("Initializing DataBase Connections ... ")
	initState := DBManager.InitRTXCollections()
	if initState {
		fmt.Println("[OK]")
	} else {
		fmt.Println("[FAILED]")
		return
	}

	fmt.Print("Initializing the server ... ")
	app := fiber.New()
	app.Use(cors.New())
	//app.Use(Middlewares.Auth)
	app.Static("/Resources", "./Resources")
	app.Use(pprof.New())

	SetupRoutes(app)
	fmt.Println("[OK]")

	app.Listen(":2022")
}
