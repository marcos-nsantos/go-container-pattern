package installment

import (
	"gorm.io/gorm"

	"github.com/marcos-nsantos/go-container-pattern/internal/structs"
)

type Repository struct {
	db *gorm.DB
}

func NewInstallmentRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(installment structs.Installment) error {
	return r.db.Create(installment).Error
}
