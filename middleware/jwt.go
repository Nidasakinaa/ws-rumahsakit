package middleware

import (
    "context"
    "os"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "google.golang.org/api/idtoken"
)

func AuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Missing token",
            })
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Invalid token format",
            })
        }

        tokenString := parts[1]


        if strings.Contains(tokenString, ".") { 
            token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
                }
                secretKey := os.Getenv("JWT_SECRET")
                return []byte(secretKey), nil
            })
            if err != nil || !token.Valid {
                return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                    "message": "Invalid JWT token",
                })
            }

            claims, ok := token.Claims.(jwt.MapClaims)
            if !ok || !token.Valid {
                return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                    "message": "Invalid token claims",
                })
            }

            adminID, ok := claims["admin_id"].(string)
            if !ok {
                return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                    "message": "Invalid token claims",
                })
            }

            role, ok := claims["role"].(string)
            if !ok || role != "admin" {
                return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                    "message": "Access denied: Admin role required",
                })
            }

            c.Locals("admin_id", adminID)
        } else { // Proses Google OAuth token
            ctx := context.Background()
            payload, err := idtoken.Validate(ctx, tokenString, os.Getenv("GOOGLE_CLIENT_ID"))
            if err != nil {
                return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                    "message": "Invalid Google OAuth token",
                })
            }

            googleUserID := payload.Subject
            role, ok := payload.Claims["role"].(string)
            if !ok || role != "admin" {
                return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                    "message": "Access denied: Admin role required",
                })
            }

            c.Locals("admin_id", googleUserID)
        }

        return c.Next()
    }
}
