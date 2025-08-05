package grpc

import (
	"context"
	"fmt"

	"github.com/Maxcillius/GoChat/platforms/db/db"
	"github.com/Maxcillius/GoChat/platforms/db/proto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type server struct {
	proto.UnimplementedDBServiceServer
	db db.DB
}

func (s *server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	user, err := s.db.CreateUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, fmt.Errorf("error while creating user: %w", err)
	}

	return &proto.CreateUserResponse{
		Id:    user.ID.String(),
		Email: user.Email,
	}, nil
}

func (s *server) CreateProfile(ctx context.Context, req *proto.CreateProfileRequest) (*proto.CreateProfileResponse, error) {
	id, err := uuid.Parse(req.Id)

	if err != nil {
		return nil, fmt.Errorf("invalid uuid: %w", err)
	}

	if len(req.DisplayName) < 3 {
		return nil, fmt.Errorf("name length should be greater than 1")
	} else if len(req.DisplayName) > 20 {
		return nil, fmt.Errorf("name length should be less than 20")
	}

	newProfile := s.db.CreateProfile(ctx, id, req.DisplayName, req.AvatarUrl, req.Bio)
	if newProfile != nil {
		return nil, fmt.Errorf("error while creating profile: %w", err)
	}

	return nil, nil
}

func (s *server) CreateSession(ctx context.Context, req *proto.CreateSessionRequest) (*proto.CreateSessionResponse, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("user id is invalid: %w", err)
	}

	userid := pgtype.UUID{Bytes: id, Valid: true}
	ipaddress := req.IpAddress
	useragent := req.UserAgent
	// auth_token := req.AuthToken

	// send auth token to the session provider server to get the refresh and access token
	access_token := "access_token"
	refresh_token := "refresh_token"

	if req.IpAddress == "" {
		return nil, fmt.Errorf("ipaddress is empty: %w", err)
	}

	if refresh_token == "" {
		return nil, fmt.Errorf("authorization token is empty: %w", err)
	}

	newSession, err := s.db.CreateSession(ctx, userid, refresh_token, access_token, ipaddress, useragent)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	time := newSession.ExpiresAt.Time.GoString()

	return &proto.CreateSessionResponse{
		Id:           newSession.ID.URN(),
		UserId:       newSession.UserID.String(),
		RefreshToken: newSession.RefreshToken,
		ExpiresAt:    time,
	}, nil
}
