// Copyright 2021 Monoskope Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package queryhandler

import (
	"context"
	"time"

	api "github.com/finleap-connect/monoskope/pkg/api/domain"
	"github.com/finleap-connect/monoskope/pkg/api/domain/projections"
	"github.com/finleap-connect/monoskope/pkg/domain/errors"
	"github.com/finleap-connect/monoskope/pkg/domain/repositories"
	grpcUtil "github.com/finleap-connect/monoskope/pkg/grpc"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// UserServer is the implementation of the TenantService API
type UserServer struct {
	api.UnimplementedUserServer

	repo repositories.ReadOnlyUserRepository
}

// NewUserServer returns a new configured instance of UserServer
func NewUserServer(userRepo repositories.ReadOnlyUserRepository) *UserServer {
	return &UserServer{
		repo: userRepo,
	}
}

func NewUserClient(ctx context.Context, queryHandlerAddr string) (*grpc.ClientConn, api.UserClient, error) {
	conn, err := grpcUtil.
		NewGrpcConnectionFactoryWithDefaults(queryHandlerAddr).
		ConnectWithTimeout(ctx, 10*time.Second)
	if err != nil {
		return nil, nil, errors.TranslateToGrpcError(err)
	}

	return conn, api.NewUserClient(conn), nil
}

// GetById returns the user found by the given id.
func (s *UserServer) GetById(ctx context.Context, userId *wrappers.StringValue) (*projections.User, error) {
	uuid, err := uuid.Parse(userId.Value)
	if err != nil {
		return nil, errors.TranslateToGrpcError(err)
	}

	user, err := s.repo.ByUserId(ctx, uuid)
	if err != nil {
		return nil, errors.TranslateToGrpcError(err)
	}
	return user.Proto(), nil
}

// GetByEmail returns the user found by the given email address.
func (s *UserServer) GetByEmail(ctx context.Context, email *wrappers.StringValue) (*projections.User, error) {
	user, err := s.repo.ByEmail(ctx, email.GetValue())
	if err != nil {
		return nil, errors.TranslateToGrpcError(err)
	}
	return user.Proto(), nil
}

func (s *UserServer) GetRoleBindingsById(userId *wrappers.StringValue, stream api.User_GetRoleBindingsByIdServer) error {
	uuid, err := uuid.Parse(userId.Value)
	if err != nil {
		return errors.TranslateToGrpcError(err)
	}

	user, err := s.repo.ByUserId(stream.Context(), uuid)
	if err != nil {
		return errors.TranslateToGrpcError(err)
	}

	for _, role := range user.Roles {
		err := stream.Send(role)
		if err != nil {
			return errors.TranslateToGrpcError(err)
		}
	}
	return nil
}

func (s *UserServer) GetAll(request *api.GetAllRequest, stream api.User_GetAllServer) error {
	users, err := s.repo.GetAll(stream.Context(), request.GetIncludeDeleted())
	if err != nil {
		return errors.TranslateToGrpcError(err)
	}

	for _, user := range users {
		err := stream.Send(user.Proto())
		if err != nil {
			return errors.TranslateToGrpcError(err)
		}
	}
	return nil
}
