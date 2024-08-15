package routes

import (
	"github.com/amrimuf/hompimRent/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userCtrl *controllers.UserController) {
    router.Post("/", userCtrl.CreateUser)
    router.Get("/", userCtrl.GetAllUsers)
    router.Get("/:id", userCtrl.GetUserByID)
    router.Put("/:id", userCtrl.UpdateUser)
    router.Delete("/:id", userCtrl.DeleteUser)
}
