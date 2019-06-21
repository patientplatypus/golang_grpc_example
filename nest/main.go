
package main

import (
	// "context"
	"log"
	// "net"
	"net/http"
	"flag"
	"fmt"
	// "reflect"

	"google.golang.org/grpc"
	// "github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	"secretsquirrel_nest/request"
	pb "secretsquirrel_nest/protomain"

)

const (
	port = ":8000"
	serverPort = ":8080"
)

var(
	serverAddr         = flag.String("server_addr", "127.0.0.1:8080", "The server address in the format of host:port")
	serverAddrDocker   = flag.String("server_addr_docker", "http://secretsquirrel_nut:8080", "The server address in the format of host:port")
)

// server is used to implement helloworld.GreeterServer.


// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.Name)
// 	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
// }

// func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//   return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
// }


func main() {

	fmt.Println("inside the nest main()")
	conn, err := grpc.Dial(serverPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewNutsClient(conn)

	r := mux.NewRouter()
	fmt.Println("after defining NewRouter()")
	fmt.Println(r)
	r.HandleFunc("/nut/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside /nut/post")
		request.PostNut(w, r, client)
	}).Methods("POST")
	r.HandleFunc("/nut/get", func(w http.ResponseWriter, r *http.Request) {
		request.GetNut(w, r, client)
	}).Methods("GET")
	r.HandleFunc("/nut/stream", func(w http.ResponseWriter, r *http.Request) {
		request.StreamNut(w, r, client)
	}).Methods("POST")

	// // CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// serve hot and fresh 15ms or less
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))

}