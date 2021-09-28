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

package main

import (
	"context"

	api_domain "github.com/finleap-connect/monoskope/pkg/api/domain"
	api_common "github.com/finleap-connect/monoskope/pkg/api/domain/common"
	api "github.com/finleap-connect/monoskope/pkg/api/eventsourcing"
	"github.com/finleap-connect/monoskope/pkg/domain"
	es "github.com/finleap-connect/monoskope/pkg/eventsourcing"
	"github.com/finleap-connect/monoskope/pkg/grpc"
	"github.com/finleap-connect/monoskope/pkg/logger"
	ggrpc "google.golang.org/grpc"

	"github.com/finleap-connect/monoskope/internal/commandhandler"
	"github.com/finleap-connect/monoskope/internal/common"
	"github.com/finleap-connect/monoskope/internal/eventstore"
	"github.com/finleap-connect/monoskope/internal/queryhandler"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
)

var (
	apiAddr          string
	metricsAddr      string
	keepAlive        bool
	eventStoreAddr   string
	queryHandlerAddr string
)

var serverCmd = &cobra.Command{
	Use:   "server [flags]",
	Short: "Starts the server",
	Long:  `Starts the gRPC API and metrics server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		ctx := context.Background()
		log := logger.WithName("server-cmd")

		// Create EventStore client
		log.Info("Connecting event store...", "eventStoreAddr", eventStoreAddr)
		conn, esClient, err := eventstore.NewEventStoreClient(ctx, eventStoreAddr)
		if err != nil {
			return err
		}
		defer conn.Close()

		// Create UserService client
		log.Info("Connecting query handler...", "queryHandlerAddr", queryHandlerAddr)
		conn, userSvcClient, err := queryhandler.NewUserClient(ctx, queryHandlerAddr)
		if err != nil {
			return err
		}
		defer conn.Close()

		// Setup domain
		log.Info("Seting up es/cqrs...")
		err = domain.SetupCommandHandlerDomain(ctx, userSvcClient, esClient)
		if err != nil {
			return err
		}

		// Create gRPC server and register implementation
		log.Info("Creating gRPC server...")
		grpcServer := grpc.NewServer("commandhandler-grpc", keepAlive)

		commandHandlerApiServer := commandhandler.NewApiServer(es.DefaultCommandRegistry)
		grpcServer.RegisterService(func(s ggrpc.ServiceRegistrar) {
			api.RegisterCommandHandlerServer(s, commandHandlerApiServer)
			api_domain.RegisterCommandHandlerExtensionsServer(s, commandHandlerApiServer)
			api_common.RegisterServiceInformationServiceServer(s, common.NewServiceInformationService())
		})

		// Finally start the server
		log.Info("gRPC server start serving...")
		return grpcServer.Serve(apiAddr, metricsAddr)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	// Local flags
	flags := serverCmd.Flags()
	flags.BoolVar(&keepAlive, "keep-alive", false, "If enabled, gRPC will use keepalive and allow long lasting connections")
	flags.StringVarP(&apiAddr, "api-addr", "a", ":8080", "Address the gRPC service will listen on")
	flags.StringVar(&metricsAddr, "metrics-addr", ":9102", "Address the metrics http service will listen on")
	flags.StringVar(&eventStoreAddr, "event-store-api-addr", ":8081", "Address the eventstore gRPC service is listening on")
	flags.StringVar(&queryHandlerAddr, "query-handler-api-addr", ":8081", "Address the queryhandler gRPC service is listening on")
}
