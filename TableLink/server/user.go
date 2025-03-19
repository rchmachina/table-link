package server

import (
	"context"
	models "tablelink/dto"
	"tablelink/proto/users"
	"tablelink/repository"
)

type UserServer struct {
	users.UnimplementedUserServiceServer
	repo repository.UserRepository
}

func NewUserServer(repo repository.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) GetAllUsers(ctx context.Context, req *users.Empty) (*users.UserListResponse, error) {
	usersData, _ := s.repo.GetAllUsers()
	var userList []*users.User

	for _, u := range usersData {
		userList = append(userList, &users.User{
			RoleId: u.RoleID, RoleName: u.RoleName, Name: u.Name, Email: u.Email, LastAccess: u.LastAccess,
		})
	}

	return &users.UserListResponse{Status: true, Message: "Successfully", Users: userList}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.UserResponse, error) {
	user := models.CreateNewUser{
		RoleID: req.RoleId, Name: req.Name, Email: req.Email, Password: req.Password,
	}
	_ = s.repo.CreateUser(user)
	return &users.UserResponse{Status: true, Message: "Successfully"}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *users.UpdateUserRequest) (*users.UserResponse, error) {
	_ = s.repo.UpdateUser(req.UserId, req.Name)
	return &users.UserResponse{Status: true, Message: "Successfully"}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (*users.UserResponse, error) {
	_ = s.repo.DeleteUser(req.UserId)
	return &users.UserResponse{Status: true, Message: "Successfully"}, nil
}
