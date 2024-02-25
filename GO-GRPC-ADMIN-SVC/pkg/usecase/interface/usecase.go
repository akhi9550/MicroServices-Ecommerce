package interfaces

import (
	"admin/service/pkg/domain"
	"admin/service/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
}
