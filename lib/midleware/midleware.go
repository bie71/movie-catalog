package midleware

import (
	"movie-catalog/config"
	"movie-catalog/lib/auth"
	"strings"

	"movie-catalog/internal/adapter/handler/response"

	"github.com/gofiber/fiber/v2"
)

type Midleware interface {
	CheckToken() fiber.Handler
}

// CheckToken implements the Midleware interface
func (o *Options) CheckToken() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var errResponse response.ErrorResponseDefault
		authHandler := c.Get("Authorization")

		if authHandler == "" {
			errResponse.Meta.Status = false
			errResponse.Meta.Message = "Authorization header is missing"
			return c.Status(fiber.StatusUnauthorized).JSON(errResponse)
		}
		tokenStr := strings.Split(authHandler, "Bearer ")[1]
		claims, err := o.authJwt.VerifyAccessToken(tokenStr)
		if err != nil {
			errResponse.Meta.Status = false
			errResponse.Meta.Message = "Invalid token"
			return c.Status(fiber.StatusUnauthorized).JSON(errResponse)
		}
		c.Locals("user", claims)
		return c.Next()
	}
}

type Options struct {
	authJwt auth.JWT
}

func NewMidleware(cfg *config.Config) Midleware {
	opt := new(Options)
	opt.authJwt = auth.NewJwt(cfg)
	return opt
}
