package routes

import (
    "github.com/amrimuf/hompimRent/controllers"
    "github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userCtrl *controllers.UserController) {
    api := app.Group("/users")
    api.Post("/", userCtrl.CreateUser)
    api.Get("/", userCtrl.GetAllUsers)
    api.Get("/:id", userCtrl.GetUserByID)
    api.Put("/:id", userCtrl.UpdateUser)
    api.Delete("/:id", userCtrl.DeleteUser)
}
