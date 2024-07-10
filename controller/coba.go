package controller

import (
	cek "github.com/Nidasakinaa/BeRS/module"
	"github.com/Nidasakinaa/ws-rumahsakit/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPasien(c *fiber.Ctx) error {
	ps := cek.GetAllPasien(config.Ulbimongoconn, "DataPasien")
	return c.JSON(ps)
}
