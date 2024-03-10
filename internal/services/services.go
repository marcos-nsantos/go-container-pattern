package services

import (
	"github.com/marcos-nsantos/go-container-pattern/internal/interfaces"
	"github.com/marcos-nsantos/go-container-pattern/internal/repositories"
	"github.com/marcos-nsantos/go-container-pattern/internal/services/installment"
)

type Container struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(repositoryContainer repositories.Container) Container {
	return Container{
		InstallmentService: installment.NewInstallmentService(repositoryContainer),
	}
}
