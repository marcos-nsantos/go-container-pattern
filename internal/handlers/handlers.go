package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marcos-nsantos/go-container-pattern/internal/handlers/installment"
	"github.com/marcos-nsantos/go-container-pattern/internal/services"
)

func NewHandlerContainer(router fiber.Router, serviceContainer services.Container) {
	installment.NewInstallmentHandler(router, serviceContainer).SetRoutes()
}
