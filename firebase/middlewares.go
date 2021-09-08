package firebase

import (
	"context"
	"fmt"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
)

type Header map[string][]string

// AuthMiddleware : to verify all authorized operations
func FirebaseAuthMiddleware(c *fiber.Ctx) error {
	firebaseAuth := c.Locals("firebaseAuth").(*auth.Client)

	authorizationToken := c.Get("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		c.JSON(map[string]string{
			"error": "invalid authentication",
		})
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(map[string]string{
			"error": "invalid token",
		})
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Set("UUID", token.UID)
	return c.Next()
}
