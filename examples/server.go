package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/clsacramento/cloudrequests"
)

func hello(w http.ResponseWriter, r *http.Request) {
	resp :=  "Hello! "+r.URL.Path[1:]

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
        fmt.Println("Body data: "+string(body))

	w.Write([]byte(resp+" "+string(body)))
}

func main() {
	http.HandleFunc("/helo/", hello)
	http.HandleFunc("/end/", cloudrequests.CollectEndpoint)
	http.HandleFunc("/listends/", cloudrequests.CollectEndpointList)
	fmt.Println("Listening starts on port 8080:")
	http.ListenAndServe(":8088", nil)
}
