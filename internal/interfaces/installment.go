package interfaces

import "github.com/marcos-nsantos/go-container-pattern/internal/structs"

type InstallmentService interface {
	Create(installment structs.Installment) error
}

type InstallmentRepository interface {
	Create(installment structs.Installment) error
}
