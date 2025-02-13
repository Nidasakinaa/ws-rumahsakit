package handler

import (
	"fmt"
	"net/http"
	"strings"

	iniconfig "github.com/Nidasakinaa/BeRS/config"
	inimodel "github.com/Nidasakinaa/BeRS/model"
	cek "github.com/Nidasakinaa/BeRS/module"
	"github.com/Nidasakinaa/ws-rumahsakit/config"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var loginDetails inimodel.User
	if err := c.BodyParser(&loginDetails); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
		})
	}

	storedAdmin, err := cek.GetByUsername(config.Ulbimongoconn, "User", loginDetails.Username)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  http.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}

	// if storedAdmin.Role != "admin" {
	// 	return c.Status(http.StatusForbidden).JSON(fiber.Map{
	// 		"status":  http.StatusForbidden,
	// 		"message": "Access denied: only admins can log in",
	// 	})
	// }
    fmt.Printf("storedAdmin: %v\n", storedAdmin)

	if !iniconfig.CheckPasswordHash(loginDetails.Password, storedAdmin.Password) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  http.StatusUnauthorized,
			"message": "Invalid credentials",
		})
	}

	token, err := iniconfig.GenerateJWT(*storedAdmin) 
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Could not generate token",
		})
	}

	err = cek.SaveTokenToDatabase(config.Ulbimongoconn, "tokens", storedAdmin.ID.Hex(), token)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Could not save token",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Login successful",
		"token":   token,
	})
}

func CustomerLogin(c *fiber.Ctx) error {
    var loginDetails inimodel.User
    if err := c.BodyParser(&loginDetails); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "status":  http.StatusBadRequest,
            "message": "Invalid request",
        })
    }

    storedUser, err := cek.GetByUsername(config.Ulbimongoconn, "User", loginDetails.Username)
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
            "status":  http.StatusUnauthorized,
            "message": "Invalid credentials",
        })
    }

    // if storedUser.Role != "customer" {
    //     return c.Status(http.StatusForbidden).JSON(fiber.Map{
    //         "status":  http.StatusForbidden,
    //         "message": "Access denied: only customers can log in",
    //     })
    // }

    if !iniconfig.CheckPasswordHash(loginDetails.Password, storedUser.Password) {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
            "status":  http.StatusUnauthorized,
            "message": "Invalid credentials",
        })
    }

    token, err := iniconfig.GenerateJWT(*storedUser) 
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Could not generate token",
        })
    }

    err = cek.SaveTokenToDatabase(config.Ulbimongoconn, "tokens", storedUser.ID.Hex(), token)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Could not save token",
        })
    }

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "status":  http.StatusOK,
        "message": "Customer login successful",
        "token":   token,
    })
}

func Logout(c *fiber.Ctx) error {
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

	token := parts[1]

	err := cek.DeleteTokenFromMongoDB(config.Ulbimongoconn, "tokens", token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete token",
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}

//Register Function

func Register(c *fiber.Ctx) error {
    var newAdmin inimodel.User
    if err := c.BodyParser(&newAdmin); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "status":  http.StatusBadRequest,
            "message": "Invalid request body",
        })
    }

    // Cek apakah username sudah ada di database
    existingUser, err := cek.GetByUsername(config.Ulbimongoconn, "User", newAdmin.Username)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Could not check existing username",
        })
    }

    if existingUser != nil {
        return c.Status(http.StatusConflict).JSON(fiber.Map{
            "status":  http.StatusConflict,
            "message": "Username already taken",
        })
    }

    // Hash password sebelum disimpan
    hashedPassword, err := iniconfig.HashPassword(newAdmin.Password)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Could not hash password",
        })
    }
    newAdmin.Password = hashedPassword

    // Simpan user baru ke database
    insertedID, err := cek.InsertUsers(config.Ulbimongoconn, "User", newAdmin.FullName, newAdmin.Phone, newAdmin.Username, newAdmin.Password)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Could not register user",
        })
    }

    return c.Status(http.StatusCreated).JSON(fiber.Map{
        "status":  http.StatusCreated,
        "message": "Account registered successfully",
        "data": fiber.Map{
            "user_id": insertedID,
        },
    })
}

func DashboardPage(c *fiber.Ctx) error {
    adminID := c.Locals("admin_id")
    if adminID == nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status":  http.StatusInternalServerError,
            "message": "Admin ID not found in context",
        })
    }

    adminIDStr := fmt.Sprintf("%v", adminID)

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "status":  http.StatusOK,
        "message": "Dashboard access successful",
        "admin_id": adminIDStr,
    })
}	
