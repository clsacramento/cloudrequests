package cloudrequests

import(
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
//	"github.com/clsacramento/cloudrequests"
)

func getRequestBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
        fmt.Println("Body data: "+string(body))

        //errors
        if err != nil {
                panic(err)
        }
        if err := r.Body.Close(); err != nil {
                panic(err)
        }	

	return body
}



func CollectEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("collect endpoint")
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	body := getRequestBody(r)

	endpoint, err := GetEndpointFromJSON(body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Endpoint decoded: ",endpoint.Url, " endpoint data: "+endpoint.Data)
	_,reason := endpoint.EndpointCheck()	

	w.Write([]byte(reason))
}

func CollectEndpointList(w http.ResponseWriter, r *http.Request) {
        fmt.Println("collect endpoint list")
        // Make sure we can only be called with an HTTP POST request.
        if r.Method != "POST" {
                w.Header().Set("Allow", "POST")
                w.WriteHeader(http.StatusMethodNotAllowed)
                return
        }

        body := getRequestBody(r)

        endpoints, err := GetEndpointListFromJSON(body)
        if err != nil {
                panic(err)
        }
	for _,endpoint := range endpoints {
	        fmt.Println("Endpoint decoded: ",endpoint.Url, " endpoint data: "+endpoint.Data)
	        _,reason := endpoint.EndpointCheck()
        	w.Write([]byte(reason))
	}
}
