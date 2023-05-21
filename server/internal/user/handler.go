package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type FiberHandler struct {
	service Service
}

func NewFiberHandler(service Service) *FiberHandler {
	return &FiberHandler{service: service}
}

func (h *FiberHandler) Create(c *fiber.Ctx) error {
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

func (h *FiberHandler) Login(c *fiber.Ctx) error {
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

func (h *FiberHandler) Logout(c *fiber.Ctx) error {
	cookie := &fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(cookie)
	return c.JSON(nil)
}

func (h *FiberHandler) TokenCheck(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	return c.JSON(claims)
}
