package main

import(
	"bytes"
//	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"github.com/clsacramento/cloudrequests"
//	"github.com/Jeffail/gabs"
)


func test_end(ep cloudrequests.Endpoint, c chan string) {
	_,reason := ep.EndpointCheck()
	c <- reason
}

func print_test(c chan string) {
	reason := <- c
	fmt.Println(reason)
}


func test_ends(ends []cloudrequests.Endpoint,c chan string) {
        for _,end := range ends {
		end.Timeout = 10000
		go test_end(end,c)
	}
}
func print_tests(c chan string) {
	for reason := range c {
		fmt.Println(reason)
	}
}

func bunktest(ends []cloudrequests.Endpoint){
      for _,end := range ends{
              _,reason := end.EndpointCheck()
//              if(success){
//                      fmt.Println("OK: ",end.Url)
//              } else {
                fmt.Println(reason)
//              }
      }	
}

func concurrenttest(ends []cloudrequests.Endpoint) {
	c := make(chan string)

	test_ends(ends,c)
//	print_tests(c)
	for i := 0; i < len(ends); i++ {
		print_test(c)
	}
}


func main() {
//	jsonParsed,_ := gabs.ParseJSON([]byte(`
//		{"auth":{"passwordCredentials":{"username": "admin", "password": "V21lxTa05JaX"},"tenantName": "demo"}}
//	`))

//	var value string
//	value,_ = jsonParsed.Path("auth.passwordCredentials.username").Data().(string)

//	fmt.Println(value)

	ends,err := cloudrequests.GetEndpointsFromFile("endpoints.json")

	if err != nil {
		panic(err)
	}

//	endchan := make(chan []cloudrequests.Endpoint)
	for i := 0; i < 1 ; i++ {
//		go func (){
		ends = append(ends,ends...)
//			endchan <- ends
//		}()
//		concurrenttest(<-endchan)
//		bunktest(<-endchan)
	}

//	concurrenttest(ends)
//	bunktest(ends)

	fmt.Println("size: ",len(ends))

	mJson := []byte(`{"auth":{"passwordCredentials":{"username": "admin", "password": "V21lxTa05JaX"},"tenantName": "demo"}}`)
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest("POST", "http://172.16.90.2:5000/v2.0/tokens", contentReader)
	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Notes","GoRequest is coming!")
	client := &http.Client{
		Timeout: 1000 * time.Millisecond,
	}
	resp, err := client.Do(req)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

	var ep cloudrequests.Endpoint 
	ep.Url =  "http://172.16.90.2:5000/v2.0/tokens"
	ep.Method = "post"
	ep.Data = `{"auth":{"passwordCredentials":{"username": "admin", "password": "V21lxTa05JaX"},"tenantName": "demo"}}` 
	request(ep)
}

func request(ep cloudrequests.Endpoint) {
	mJson := []byte(ep.Data)
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest(strings.ToUpper(ep.Method), ep.Url, contentReader)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(string(body))
}
