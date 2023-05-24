package user

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	Service
}

func NewAuthMiddleware(userService Service) *AuthMiddleware {
	return &AuthMiddleware{userService}
}

func (m *AuthMiddleware) Authenticate(c *fiber.Ctx) error {
	if c.Path() == "/api/login" || c.Path() == "/api/register" || c.Path() == "/api/logout" {
		return c.Next()
	}
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("claims", claims)
	return c.Next()
}
