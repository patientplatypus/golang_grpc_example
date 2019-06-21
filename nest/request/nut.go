package request

import(
	"context"
	"fmt"
	"time"
	"net/http"
	"log"
	"io"
	"encoding/json"
	"secretsquirrel_nest/response"
	pb "secretsquirrel_nest/protomain"
)

type StreamPost struct{
	Nuts string
}

type NutPost struct{
	Nuts string
}

func PostNut(w http.ResponseWriter, r *http.Request, client pb.NutsClient) {
	fmt.Println("inside PostNut")
	fmt.Println("and value of client: ")
	fmt.Println(client)

	decoder := json.NewDecoder(r.Body)
	var nutPost NutPost
	err := decoder.Decode(&nutPost)
	if err != nil {
		fmt.Println(err.Error())
		response.ERRORresponse(w, r, err.Error())
	}
	fmt.Println("value of nutPost")
	fmt.Println(nutPost)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	nutReturn, err := client.NutServe(ctx, &pb.NutMessage{Message: nutPost.Nuts})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	fmt.Println("after nutReturn and value")
	fmt.Println(nutReturn)
}

func GetNut(w http.ResponseWriter, r *http.Request, client pb.NutsClient) {
	fmt.Println("inside Get	Nut")
	fmt.Println("and value of client: ")
	fmt.Println(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	nutReturn, err := client.NutServe(ctx, &pb.NutMessage{Message:"get_nut_request"})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(nutReturn)
}

func StreamNut(w http.ResponseWriter, r *http.Request, client pb.NutsClient) {
	fmt.Println("inside StreamNut")
	fmt.Println("and value of client: ")
	fmt.Println(client)

	decoder := json.NewDecoder(r.Body)
	var streamPost StreamPost
	err := decoder.Decode(&streamPost)
	if err != nil {
		fmt.Println(err.Error())
		response.ERRORresponse(w, r, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.BiDiServe(ctx)
	if err != nil {
		log.Fatalf("%v.BiDiServe(_) = _, %v", client, err)
	}

	msg := pb.BiDiMessage{Nuts: streamPost.Nuts}
	if err := stream.Send(&msg); err != nil {
		fmt.Println(err)
	}
	reply, err := stream.Recv()
	if err != nil {
		log.Fatalf("%v.stream.Recv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a nut : %v", err)
			}
			log.Printf("Got nut", in.Nuts)
		}
	}()
	<-waitc
}