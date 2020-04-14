package rpc

import (
	"fmt"
	"github.com/benka-me/laruche-api/go-pkg/larapi"
	"github.com/benka-me/laruche-server/go-pkg/larsrv"
	"github.com/benka-me/laruche/go-pkg/discover"
	"github.com/benka-me/users/go-pkg/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {
	Clients
}

var grpcServer *grpc.Server

func Server_2_0(engine discover.Engine) {
	var err error
	port, err := engine.ThisPort("benka-me/laruche-api")
	if err != nil {
		log.Fatal(err)
	}
	app := &App{
		Clients: InitClients(engine, grpc.WithInsecure()), // Init clients of dependencies services
	}

	grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	{
		larapi.RegisterLarapiServer(grpcServer, app) // Register your service server.
		larsrv.RegisterLarsrvServer(grpcServer, app)
		users.RegisterUsersServer(grpcServer, app)
		reflection.Register(grpcServer)
	}

	fmt.Println("benka-me/laruche-api service running on port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
