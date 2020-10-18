package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	pb "github.com/vladfr/gosaas/helloworld"
)

type AppConfig struct {
	TLS      *bool
	certFile *string
	keyFile  *string
	port     *int
	gRPCPort *int
	debug    *bool
}

var cfg = &AppConfig{}

func parseFlags(cfg *AppConfig) {
	cfg.TLS = flag.Bool("tls", lookupEnvOrBool("TLS", false), "Connection uses TLS if true, else plain TCP")
	cfg.certFile = flag.String("cert_file", "service.pem", "The TLS cert file")
	cfg.keyFile = flag.String("key_file", "service.key", "The TLS key file")
	cfg.port = flag.Int("port", lookupEnvOrInt("PORT", 8000), "The server port")
	cfg.gRPCPort = flag.Int("gport", lookupEnvOrInt("GPORT", 9000), "The server port")
	cfg.debug = flag.Bool("debug", lookupEnvOrBool("DEBUG", false), "Turn debug mode on")
	flag.Parse()
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func startHttp(grpcServer *grpc.Server) {
	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	)
	gRPCHandler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		log.Println("[http] request")
		wrappedGrpc.ServeHTTP(resp, req)
	})

	tlscfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	if *cfg.TLS {
		srv := &http.Server{
			Addr:         fmt.Sprintf(":%d", *cfg.port),
			Handler:      gRPCHandler,
			TLSConfig:    tlscfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		log.Fatal(srv.ListenAndServeTLS(*cfg.certFile, *cfg.keyFile))
	} else {
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", *cfg.port),
			Handler: gRPCHandler,
		}
		log.Fatal(srv.ListenAndServe())
	}
}

func main() {
	parseFlags(cfg)

	log.Printf("app.status=starting http.port=%d grpc.port=%d tls=%v", *cfg.port, *cfg.gRPCPort, *cfg.TLS)
	defer log.Println("app.status=shutdown")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *cfg.TLS {
		creds, err := credentials.NewServerTLSFromFile(*cfg.certFile, *cfg.keyFile)
		if err != nil {
			log.Fatalf("[grpc] Failed to load TLS credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	if *cfg.debug {
		reflection.Register(grpcServer)
	}
	pb.RegisterGreeterServer(grpcServer, &server{})

	go startHttp(grpcServer)
	go log.Fatal(grpcServer.Serve(lis))
	for {
	}
}
