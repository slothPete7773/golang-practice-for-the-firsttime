package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// extractUserFromJWT is a middleware that extracts user data from the JWT token
func extractUserFromJWT(c *fiber.Ctx) error {
	user := &UserData{}

	// Extract the token from the Fiber context (inserted by the JWT middleware)
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	fmt.Println(claims)

	user.Email = claims["email"].(string)
	user.Role = claims["role"].(string)

	// Store the user data in the Fiber context
	c.Locals(userContextKey, user)

	return c.Next()
}

// isAdmin checks if the user is an admin
func isAdmin(c *fiber.Ctx) error {
	user := c.Locals(userContextKey).(*UserData)

	if user.Role != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
