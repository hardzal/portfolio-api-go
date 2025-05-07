package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/config"
	"github.com/hardzal/portfolio-api-go/database"
	"github.com/hardzal/portfolio-api-go/handlers/about"
	"github.com/hardzal/portfolio-api-go/handlers/auth"
	"github.com/hardzal/portfolio-api-go/handlers/project"
	"github.com/hardzal/portfolio-api-go/handlers/stack"
	"github.com/hardzal/portfolio-api-go/handlers/work"
	"github.com/hardzal/portfolio-api-go/repositories"
	"github.com/hardzal/portfolio-api-go/routes"
	"github.com/hardzal/portfolio-api-go/services"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load to env: %v", err)
	}

	gormDB, err := database.NewPostgresDB(cfg)

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	db, err := gormDB.DB()

	if err != nil {
		log.Fatalf("Failed to call DB(): %v", err)
	}

	defer db.Close()

	database.MigrateDB(gormDB)

	// login services
	userRepo := repositories.NewUserRepository(gormDB)
	authService := services.NewAuthService(userRepo)
	authHandler := auth.NewAuthHandler(authService)

	// project services
	projectRepo := repositories.NewProjectRepository(gormDB)
	projectService := services.NewProjectService(projectRepo)
	projectHandler := project.NewProjectHandler(projectService)

	// work services
	workRepo := repositories.NewWorkRepository(gormDB)
	workService := services.NewWorkService(workRepo)
	workHandler := work.NewWorkHandler(workService)

	// stack services
	stackRepo := repositories.NewStackRepository(gormDB)
	stackService := services.NewStackService(stackRepo)
	stackHandler := stack.NewStackHandler(stackService)

	// about services
	aboutRepo := repositories.NewAboutRepository(gormDB)
	aboutService := services.NewAboutService(aboutRepo)
	aboutHandler := about.NewAboutHandler(aboutService)

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Welcome to Golang Fiber + PostgreSQL + Gorm",
		})
	})

	apiRoute := app.Group("/api")

	routes.AuthRoutes(apiRoute.Group("/auth"), authHandler)
	routes.ProjectRoutes(apiRoute.Group("/projects"), projectHandler)
	routes.WorkRoutes(apiRoute.Group("/works"), workHandler)
	routes.StackRoutes(apiRoute.Group("/stacks"), stackHandler)
	routes.AboutRoutes(apiRoute.Group("/about"), aboutHandler)

	log.Println("🚀 Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
