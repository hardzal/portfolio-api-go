package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/config"
	"github.com/hardzal/portfolio-api-go/database"
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

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		var tables []string
		gormDB.Raw(
			"SELECT table_name FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA()",
		).Scan(&tables)
		return ctx.JSON(tables)
	})

	log.Println("🚀 Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
