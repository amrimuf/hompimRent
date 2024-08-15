package bootstrap

import (
	"log"
	"time"

	"github.com/amrimuf/hompimRent/controllers"
	"github.com/amrimuf/hompimRent/database"
	"github.com/amrimuf/hompimRent/middleware"
	"github.com/amrimuf/hompimRent/repositories"
	"github.com/amrimuf/hompimRent/routes"
	"github.com/amrimuf/hompimRent/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type App struct {
    server *fiber.App
	db     *pgxpool.Pool 
}

func NewApp() *App {
    return &App{
        server: fiber.New(),
    }
}

func (a *App) Bootstrap() {
	// Initialize database connection
	db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    a.db = db

	a.server.Use(middleware.ErrorLogger())
	a.server.Use(middleware.RateLimit(500, time.Minute))

	// Initialize repositories
	listingRepo := repositories.NewListingRepository(db)
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	listingService := services.NewListingService(listingRepo)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userRepository)

	// Initialize controllers
	listingController := controllers.NewListingController(listingService)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)

	// Setup routes
	ctrl := routes.Controllers{
		ListingController: listingController,
		UserController: userController,
		AuthController: authController,
	}
	routes.SetupRoutes(a.server, ctrl)

}

func (a *App) Start(addr string) error {
    defer a.db.Close()

    return a.server.Listen(addr)
}