package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	pb "github.com/vladfr/gosaas/helloworld"
)

var (
	tls      = flag.Bool("tls", lookupEnvOrBool("TLS", false), "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "service.pem", "The TLS cert file")
	keyFile  = flag.String("key_file", "service.key", "The TLS key file")
	port     = flag.Int("port", lookupEnvOrInt("PORT", 9000), "The server port")
	debug    = flag.Bool("debug", lookupEnvOrBool("DEBUG", false), "Turn debug mode on")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()

	log.Printf("app.status=starting app.port=%d", *port)
	defer log.Println("app.status=shutdown")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	if *debug {
		reflection.Register(grpcServer)
	}
	pb.RegisterGreeterServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
