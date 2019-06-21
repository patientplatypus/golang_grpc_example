package request

import(
	"fmt"
	"net/http"
	"secretsquirrel_nest/response"
)

func PostTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("inside PostTest");
	response.MESSAGEresponse(w, r, "hello there GetTest", "200")
}
