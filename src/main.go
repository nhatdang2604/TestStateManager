package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/nhatdang2604/TestStateManager/src/constants"
	"github.com/nhatdang2604/TestStateManager/src/controllers"
	"github.com/nhatdang2604/TestStateManager/src/helpers"
	"github.com/nhatdang2604/TestStateManager/src/protos"
	"google.golang.org/grpc"
)

type Server struct {
	Constant   *constants.ConfigConstant
	Controller *controllers.TestStateManagementController
}

func NewServer() *Server {

	configConstant := constants.GetConfigConstant()

	server := &Server{
		Constant:   configConstant,
		Controller: controllers.NewTestStateManagementController(),
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

	//Trying to close the database, after the server has been shutting down
	db := helpers.GetDB()
	defer db.Connection.Close()

	go s.Controller.TestStateService.StartTest(1, 2)
	go s.Controller.TestStateService.StartTest(1, 3)

	time.Sleep(10 * time.Second)
	go s.Controller.TestStateService.GetRemainTimeOfInprogressTest(15)
	go s.Controller.TestStateService.GetRemainTimeOfInprogressTest(16)

	//Start the server
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
