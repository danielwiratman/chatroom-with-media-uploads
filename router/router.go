package router

import (
	"github.com/danielwiratman/chatroom-with-media-uploads/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupAPI(router fiber.Router, userHandler *user.UserFiberHandler) {
  router.Post("/register", userHandler.Create)
	router.Post("/login", userHandler.Login)
  router.Post("/logout", userHandler.Logout)
  router.Get("/token-check", userHandler.TokenCheck)
	router.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
}
