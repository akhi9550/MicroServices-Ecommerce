package interfaces

import "grpc-api-gateway/pkg/utils/models"

type AdminClient interface {
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)
}
