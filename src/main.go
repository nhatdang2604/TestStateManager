package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/nhatdang2604/TestStateManager/src/constants"
	"github.com/nhatdang2604/TestStateManager/src/controllers"
	"github.com/nhatdang2604/TestStateManager/src/protos"
	"google.golang.org/grpc"
)

type Server struct {
	Constant   *constants.ConfigConstant
	Controller *controllers.TestStateManagementController
}

func NewServer() *Server {

	configConstant, err := constants.NewConfigConstant()

	if nil != err {
		log.Fatalf("Error while creating server")
		return nil
	}

	server := &Server{
		Constant:   configConstant,
		Controller: new(controllers.TestStateManagementController),
	}
	return server
}

func (s *Server) Start() error {

	//Listening to the internet
	host := s.Constant.Host
	port := s.Constant.Port
	listener, err := net.Listen("tcp", strings.Join([]string{host, port}, ":"))

	if nil != err {
		log.Fatalf("Error while creating listener: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	protos.RegisterTestStateManagementServer(grpcServer, s.Controller)
	fmt.Println("Test State Manager Service is running")

	err = grpcServer.Serve(listener)

	if nil != err {
		log.Fatalf("Error while serving: %v", err)
	}

	return err
}

func main() {
	server := NewServer()
	server.Start()
}
