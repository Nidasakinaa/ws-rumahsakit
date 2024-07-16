package url

import (
	"github.com/Nidasakinaa/ws-rumahsakit/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage)
	page.Get("/pasien", controller.GetPasien)
	page.Get("/pasien/:id", controller.GetPasienID)
	page.Post("/insert", controller.InsertDataPasien)
	page.Put("/update/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePasienByID)

	page.Get("/docs/*", swagger.HandlerDefault)
}
