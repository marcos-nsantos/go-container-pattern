package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/marcos-nsantos/go-container-pattern/internal/handlers"
	"github.com/marcos-nsantos/go-container-pattern/internal/repositories"
	"github.com/marcos-nsantos/go-container-pattern/internal/services"
	"github.com/marcos-nsantos/go-container-pattern/internal/structs"
)

func main() {
	server := fiber.New()

	dialector := postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	))

	db, err := gorm.Open(dialector)
	if err != nil {
		log.Fatalf("banco não conectou. Err=%v", err.Error())
	}

	err = db.AutoMigrate(&structs.Installment{})
	if err != nil {
		log.Fatalf("erro ao executar migração. Err=%v", err)
	}

	repositoryContainer := repositories.GetRepositories(db)
	servicesContainer := services.GetServices(repositoryContainer)
	handlers.NewHandlerContainer(server, servicesContainer)

	server.Get("/heath", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("pong :)")
	})

	if err := server.Listen(":3001"); err != nil {
		log.Fatalf("aplicação não subiu. Err%v", err.Error())
	}
}
