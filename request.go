package cloudrequests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
//	"net/http"
//	"reflect"
//	"github.com/clsacramento/
//	"github.com/parnurzeal/gorequest"
)

func get_time(t int) time.Duration {
	return time.Duration(t) * time.Millisecond
}

func make_query(ep Endpoint) *http.Request {	
	contentReader := bytes.NewReader([]byte(ep.Data))
	if ep.Method == "" {
		ep.Method = "GET"
	}
	req, _ := http.NewRequest(strings.ToUpper(ep.Method), ep.Url, contentReader)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func do_request(ep Endpoint) (Response) {
	query := make_query(ep)

	client := &http.Client{
		Timeout: get_time(ep.Timeout),
	}
	resp, err := client.Do(query)

	var endresp Response
	if err != nil {
		endresp.ErrorMessage = err.Error()
		return endresp
	}
	bodybytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodybytes)

	fmt.Println(resp)
	fmt.Println(body)
	fmt.Println(err)


	endresp.Status = resp.StatusCode
	endresp.Data = body
	return endresp
}

func (ep Endpoint) EndpointCheck() (bool,string) {
	resp := do_request(ep)
	if ep.Name == ""{
		ep.Name = ep.Url
	}
	if resp.Status == 0 {
		timeout,_ := regexp.MatchString("timeout",resp.ErrorMessage)
		if timeout {
			return false, ep.Name+" is down: request timed out"
		}
		return false, ep.Name+"is down: "+resp.ErrorMessage
	} else if ep.Expected.Status != 0 && ep.Expected.Status != resp.Status {
		strstatexp := strconv.Itoa(ep.Expected.Status)
		strstatresp := strconv.Itoa(resp.Status)
		return false, ep.Name+" is up, BUT expecteded status ("+strstatexp+") did not match response status: "+strstatresp+" -msg: "+resp.Data
				//+string(resp.Status)
	} else if ep.Expected.Data != "" { 

		matched,_ := regexp.MatchString(ep.Expected.Data, resp.Data)
		if !matched {
			return false, ep.Name+" is up, BUT expected data ("+ep.Expected.Data+") did not match response data: "+resp.Data
		}
	}
	return true, ep.Name+" is up: "+ep.Url
}

//func main() {
//	ep := Endpoint{Url: "http://172.16.90.2:35357/v3", Method: "Get", Timeout: 7}
//
//	for i := 1; i <= 30; i++ {
//		if i > 3 && i < 20 {
//			ep.Expected.Status = 400
//		} else {
//			ep.Expected.Status = 200
//			ep.Expected.Data = "openstack" 
//		}
//		_,reason := EndpointCheck(ep)
//		fmt.Println(reason)
//	}
//}
