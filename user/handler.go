package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserFiberHandler struct {
	service UserService
}

func NewUserFiberHandler(service UserService) *UserFiberHandler {
	return &UserFiberHandler{service: service}
}

func (h *UserFiberHandler) Create(c *fiber.Ctx) error {
	req := &CreateUserReq{}
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, err := h.service.Create(c.Context(), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (h *UserFiberHandler) Login(c *fiber.Ctx) error {
	req := &LoginUserReq{}
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, err := h.service.Login(c.Context(), req)
	if err != nil {
		return err
	}
	cookie := &fiber.Cookie{
		Name:     "jwt",
		Value:    res.Token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(cookie)
	return c.JSON(res)
}

func (h *UserFiberHandler) Logout(c *fiber.Ctx) error {
	cookie := &fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(cookie)
	return c.JSON(nil)
}

func (h *UserFiberHandler) TokenCheck(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	return c.JSON(claims)
}
