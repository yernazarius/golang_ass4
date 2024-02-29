package internal

import (
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/use_case"
)

func NewContactRepository() repository.IRepository {
	return nil
}

func NewContactUseCase(repo repository.ContactRepository) useCase.ContactUseCase {
	return nil
}

func NewContactDelivery(useCase useCase.ContactUseCase) delivery.ContactDelivery {
	return nil
}
