package url

import (
	"github.com/Nidasakinaa/ws-rumahsakit/controller"
	"github.com/Nidasakinaa/ws-rumahsakit/handler"
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

	page.Get("/user", controller.GetAllUser)
	page.Get("/user/:id", controller.GetUserID)
	page.Post("/insertUser", controller.InsertDataUser)
	page.Put("/user/updateUser/:id", controller.UpdateDataUser)
	page.Delete("/user/deleteUser/:id", controller.DeleteUserByID)
	page.Post("/registeruser", handler.Register)
	page.Post("/login", handler.Login)

	page.Get("/docs/*", swagger.HandlerDefault)
	page.Get("/dashboard", handler.DashboardPage)
}
