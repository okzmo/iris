package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/okzmo/noname/pb"
	"github.com/okzmo/noname/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	pb.UnimplementedChatAppServer
	config util.Config
}

func newServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}

func RunGRPCServer(config util.Config) {
	server, err := newServer(config)
	if err != nil {
		log.Fatal("cannot create server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatAppServer(grpcServer, server)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("start GRPC server at %s", lis.Addr().String())
	if e := grpcServer.Serve(lis); e != nil {
		log.Fatal(err)
	}
}

func RunGRPCGateway(config util.Config) {
	server, err := newServer(config)
	if err != nil {
		log.Fatal("cannot create server")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	err = pb.RegisterChatAppHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server", err)
	}

	fmt.Println(grpcMux)

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	lis, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Printf("start HTTP gateway server at %s", lis.Addr().String())
	if e := http.Serve(lis, mux); e != nil {
		log.Fatal("cannot start HTTP gateway Server", err)
	}
}
