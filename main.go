package main

import (
	"github.com/danielwiratman/chatroom-with-media-uploads/db"
	"github.com/danielwiratman/chatroom-with-media-uploads/router"
	"github.com/danielwiratman/chatroom-with-media-uploads/user"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := db.NewDB()
	userRepository := user.NewUserRepository()
	userService := user.NewUserService(userRepository, db)
	userHandler := user.NewUserFiberHandler(userService)
	authMiddleware := user.NewAuthMiddleware(userService)

	app := fiber.New()
	app.Use(authMiddleware.Authenticate)
 
	api := app.Group("/api")
	router.SetupAPI(api, userHandler)

	app.Listen(":8080")
}
