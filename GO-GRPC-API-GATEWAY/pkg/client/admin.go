package client

import (
	"context"
	"errors"
	"fmt"
	interfaces "grpc-api-gateway/pkg/client/interface"
	"grpc-api-gateway/pkg/config"
	pb "grpc-api-gateway/pkg/pb/admin"
	"grpc-api-gateway/pkg/utils/models"

	"google.golang.org/grpc"
)

type adminClient struct {
	Client pb.AdminClient
}

func NewAdminClient(cfg config.Config) interfaces.AdminClient {

	grpcConnection, err := grpc.Dial(cfg.AdminSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewAdminClient(grpcConnection)

	return &adminClient{
		Client: grpcClient,
	}

}
func (ad *adminClient) AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminLogin(context.Background(), &pb.AdminLoginInRequest{
		Email:    adminDetails.Email,
		Password: adminDetails.Password,
	})

	if err != nil {
		return models.TokenAdmin{}, err
	}
	if admin.Error != "" {
		return models.TokenAdmin{}, errors.New(admin.Error)
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}
