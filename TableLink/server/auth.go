package server

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	r "tablelink/db/nosql"
	"tablelink/proto/auth"
	"tablelink/repository"
	"time"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	repo repository.AuthRepository
}

func NewAuthServer(repo repository.AuthRepository) *AuthServer {
	return &AuthServer{repo: repo}
}

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	_, err := s.repo.Login(req.Email, req.Password)
	if err != nil {
		return &auth.LoginResponse{Status: false, Message: err.Error()}, nil
	}

	randToken, err := GenerateRandomToken(32)
	randToken = "Bearer "+ randToken
	err = r.SetKey(fmt.Sprintf("token-username-%s", req.Email),randToken,time.Duration(10000))
	return &auth.LoginResponse{Status: true, Message: "Successfully", AccessToken: randToken}, nil
}

func GenerateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
