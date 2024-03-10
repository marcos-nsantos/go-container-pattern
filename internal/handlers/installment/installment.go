package installment

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/marcos-nsantos/go-container-pattern/internal/interfaces"
	"github.com/marcos-nsantos/go-container-pattern/internal/services"
	"github.com/marcos-nsantos/go-container-pattern/internal/structs"
)

type Handler struct {
	router             fiber.Router
	installmentService interfaces.InstallmentService
}

func NewInstallmentHandler(router fiber.Router, serviceContainer services.Container) Handler {
	return Handler{
		router:             router,
		installmentService: serviceContainer.InstallmentService,
	}
}

func (h Handler) SetRoutes() {
	group := h.router.Group("/installment")

	group.Post("", h.CreateInstallment)
}

func (h Handler) CreateInstallment(c *fiber.Ctx) error {
	var body structs.Installment

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = h.installmentService.Create(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON("installment created")
}
