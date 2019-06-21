package response

import(
	"fmt"
	"net/http"
	"encoding/json"
)

type ErrorResponse struct{
	Status string
	Message string
}

type MessageResponse struct{
	Status string
	Message string
}

func ERRORresponse(w http.ResponseWriter, req *http.Request, message string){
	fmt.Println("inside ERRORresponse")
	response := ErrorResponse{Message: message, Status: "400"}
	json.NewEncoder(w).Encode(response) 
}


func MESSAGEresponse(w http.ResponseWriter, req *http.Request, message string, status string){
	fmt.Println("inside MESSAGEresponse")
	response := MessageResponse{Message: message, Status: status}
	json.NewEncoder(w).Encode(response) 
}


