package app

import (
	"fmt"
	"log"
	"net"
	"os"

	"app/lib/database"
	"app/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func configServer() (port int, pk string, cert string) {
	port = viper.GetInt("APP_PORT")
	pk = viper.GetString("SSL_PRIVATE_PATH")
	cert = viper.GetString("SSL_CERT_PATH")
	if port == 0 || port > 65535 {
		log.Fatalf("port %d not allow", port)
	}
	if _, err := os.Stat(pk); os.IsNotExist(err) {
		log.Fatal(err)
	}
	if _, err := os.Stat(cert); os.IsNotExist(err) {
		log.Fatal(err)
	}
	return port, pk, cert
}

// RunHTTP is run http/2 server
func RunHTTP() error {
	defer database.DB().Close()
	port, pk, cert := configServer()
	r := gin.Default()
	router.Gin(r)
	return r.RunTLS(fmt.Sprintf(":%d", port), cert, pk) // listen and serve on 0.0.0.0:8080
}

// GRPC is gRPC server
func GRPC() *grpc.Server {
	gin.SetMode(gin.ReleaseMode)
	// pk, certp := configServer()
	port := viper.GetInt("APP_GRPC_PORT")
	if port == 0 || port > 65535 {
		os.Exit(1)
	}
	// cert, err := tls.LoadX509KeyPair(certp, pk)
	// if err != nil {
	// 	os.Exit(1)
	// }
	opts := []grpc.ServerOption{
		// Enable TLS for all incoming connections.
		// grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	serve := grpc.NewServer(opts...)
	router.RegisterGRPCServer(serve)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		os.Exit(1)
	}
	log.Printf("Start gRPC server on :%d", port)
	go func() {
		if err := serve.Serve(lis); err != nil {
			os.Exit(1)
		}
	}()
	return serve
}

// // RunGRPC is run gRPC server
// func RunGRPC() error {
// 	defer database.DB().Close()
// 	// port, pk, certp := configServer()
// 	port := viper.GetInt("APP_GRPC_PORT")

// 	// cert, err := tls.LoadX509KeyPair(certp, pk)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	opts := []grpc.ServerOption{
// 		// Enable TLS for all incoming connections.
// 		// grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
// 	}
// 	s := grpc.NewServer(opts...)
// 	router.RegisterGRPCServer(s)
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Start gRPC server on :%d", port)
// 	return s.Serve(lis)
// }
