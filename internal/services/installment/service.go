package installment

import (
	"errors"
	"time"

	"github.com/marcos-nsantos/go-container-pattern/internal/interfaces"
	"github.com/marcos-nsantos/go-container-pattern/internal/repositories"
	"github.com/marcos-nsantos/go-container-pattern/internal/structs"
)

type Service struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func NewInstallmentService(repositoryContainer repositories.Container) *Service {
	return &Service{
		InstallmentRepository: repositoryContainer.InstallmentRepository,
	}
}

func (s *Service) Create(installment structs.Installment) error {
	if installment.DueDay < time.Now().Day() {
		return errors.New("parcela vencida")
	}

	return s.InstallmentRepository.Create(installment)
}
