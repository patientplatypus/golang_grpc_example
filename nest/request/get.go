package request

import(
	"fmt"
	"net/http"
	"secretsquirrel_nest/response"
)

func GetTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("inside GetTest");
	response.MESSAGEresponse(w, r, "hello there GetTest", "200")
}
