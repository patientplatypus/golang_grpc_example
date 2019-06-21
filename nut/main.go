
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	// "time"
	"io"
	// "flag"
	"fmt"
	"google.golang.org/grpc"
	pb "secretsquirrel_nest/protomain"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) NutServe(ctx context.Context, in *pb.NutMessage) (*pb.NutReply) {
	log.Printf("Received: %v", in.Message)
	return &pb.NutReply{Reply: "Hello I heard you say" + in.Message}
}

func (s *server) BiDiServe(stream pb.Nuts_BiDiServeServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		resp := pb.BiDiMessage{Nuts: in.Nuts}

		if err := stream.Send(&resp); err != nil {
			return err
		}
	}
}

func main() {
	fmt.Println("inside the backend main()")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNutsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	handler := func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("value of req.body: ")
		s.ServeHTTP(resp, req)
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println("failed starting http server:");
		fmt.Printf(err.Error())
	}
}