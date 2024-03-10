package repositories

import (
	"gorm.io/gorm"

	"github.com/marcos-nsantos/go-container-pattern/internal/interfaces"
	"github.com/marcos-nsantos/go-container-pattern/internal/repositories/installment"
)

type Container struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func GetRepositories(db *gorm.DB) Container {
	return Container{
		InstallmentRepository: installment.NewInstallmentRepository(db),
	}
}
