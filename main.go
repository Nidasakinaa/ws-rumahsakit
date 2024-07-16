package main

import (
	"log"

	"github.com/Nidasakinaa/ws-rumahsakit/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/Nidasakinaa/ws-rumahsakit/url"

	"github.com/Nidasakinaa/ws-rumahsakit/docs"
	"github.com/gofiber/fiber/v2"
)

// @title TES SWAGGER PASIEN
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/Nidasakinaa
// @contact.email 714220040@std.ulbi.ac.id

// @host ws-rumahsakit-2e1eb71f14e2.herokuapp.com/
// @BasePath /
// @schemes https http
func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
