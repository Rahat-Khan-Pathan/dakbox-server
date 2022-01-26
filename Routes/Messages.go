package Routes

import (
	"example.com/seen-tech-rtx/Controllers"
	"github.com/gofiber/fiber/v2"
)

func BranchRoute(route fiber.Router) {
	route.Post("/new", Controllers.MessagesNew)
	route.Post("/get_all", Controllers.MessagesGetAll)
}
