package service

import (
	"admin/service/pkg/pb"
	interfaces "admin/service/pkg/usecase/interface"
	"admin/service/pkg/utils/models"
	"context"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminServer
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AdminServer {

	return &AdminServer{
		adminUseCase: useCase,
	}

}
func (ad *AdminServer) AdminLogin(ctx context.Context, Req *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {
	adminLogin := models.AdminLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	admin, err := ad.adminUseCase.LoginHandler(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{
			Error: err.Error(),
		}, err
	}
	adminDetails := &pb.AdminDetails{Id: uint64(admin.Admin.ID), Firstname: admin.Admin.Firstname, Lastname: admin.Admin.Lastname, Email: admin.Admin.Email}
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}
