package service

import (
	"context"
	"user/service/pkg/pb"
	interfaces "user/service/pkg/usecase/interface"
	"user/service/pkg/utils/models"
)

type UserServer struct {
	userUseCase interfaces.UserUseCase
	pb.UnimplementedUserServer
}

func NewUserServer(useCase interfaces.UserUseCase) pb.UserServer {

	return &UserServer{
		userUseCase: useCase,
	}

}
func (a *UserServer) UserSignUp(ctx context.Context, userSignUpDetails *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {

	userCreateDetails := models.UserSignUp{
		Firstname: userSignUpDetails.Firstname,
		Lastname:  userSignUpDetails.Lastname,
		Email:     userSignUpDetails.Email,
		Phone:     userSignUpDetails.Phone,
		Password:  userSignUpDetails.Password,
	}

	data, err := a.userUseCase.UsersSignUp(userCreateDetails)
	if err != nil {
		return &pb.UserSignUpResponse{}, err
	}
	userDetails := &pb.UserDetails{Id: uint64(data.User.ID), Firstname: data.User.Firstname, Lastname: data.User.Firstname, Email: data.User.Email, Phone: data.User.Phone}
	return &pb.UserSignUpResponse{
		Status:       200,
		UserDetails:  userDetails,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}, nil

}
