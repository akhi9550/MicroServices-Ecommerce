package interfaces

import (
	"admin/service/pkg/domain"
	"admin/service/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
}
