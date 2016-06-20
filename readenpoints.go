package cloudrequests

import(
	"fmt"
//	"encoding/json"
	"io/ioutil"
	"github.com/Jeffail/gabs"
//	"github.com/clsacramento"
)

type Endpoints struct {
	endpoints []Endpoint
}

func GetEndpointsFromFile(filename string) ([]Endpoint,error) {
	jsonstr,err := ioutil.ReadFile(filename)

	ends, err := GetEndpointListFromJSON(jsonstr)

	return ends, err
}

func GetEndpointListFromJSON(jsonbytes []byte) ([]Endpoint,error) {
	ends := make([]Endpoint,0)
	jsonParsed, err := gabs.ParseJSON(jsonbytes)
        if err != nil {
                return nil,err
        }
	children, _ := jsonParsed.S("endpoints").Children()
	for _, child := range children {
		end := parseEndpoint(child)
		ends = append(ends,end)
	}	
	return ends, err
}

func GetEndpointFromJSON(jsonbytes []byte) (Endpoint,error){
	var end Endpoint
	jsonParsed, err := gabs.ParseJSON(jsonbytes)
	if err != nil {
		return end,err
	}
	end = parseEndpoint(jsonParsed)
	return end,err 
}

func parseEndpoint(jsonParsed *gabs.Container) Endpoint {
	var end Endpoint
	url, ok := jsonParsed.Path("url").Data().(string)
	fmt.Println(url)
	if ok {
		end.Url = url
	}
	name, ok := jsonParsed.Path("name").Data().(string)
	if ok {
		end.Name = name
	}
	header, ok := jsonParsed.Path("header").Data().(string)
	if ok {
		end.Header = header
	}
	data := jsonParsed.Path("data").String()
	fmt.Println("what "+data)
	end.Data = data

	method, ok := jsonParsed.Path("method").Data().(string)
	if ok {
		end.Method = method
	}
	proxy, ok := jsonParsed.Path("proxy").Data().(string)
	if ok {
		end.Proxy = proxy
	}
	timeout, ok := jsonParsed.Path("timeout").Data().(float64)
	if ok {
		end.Timeout = int(timeout)
	}
	expected_status, ok := jsonParsed.Path("expected.status").Data().(float64)
	if ok {
		end.Expected.Status = int(expected_status)
	}
	expected_header, ok := jsonParsed.Path("expected.header").Data().(string)
	if ok {
		end.Expected.Header = expected_header
	}
	expected_data, ok := jsonParsed.Path("expected.data").Data().(string)
	if ok {
		end.Expected.Data = expected_data
	}	
	
	return end
}

//func main() {
//	jsonstr := []byte(`[{"url":"16.172.92.2:5000","timeout":10}]`)
//	jsonstr,fileerr := ioutil.ReadFile("/home/stack/gowork/src/github.com/clsacramento/lib/endpoints.json")
//	ends,err := GetEndpointsFromFile("endpoints.json")
//
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(ends)
//	for _,end := range ends {
//		_,reason := EndpointCheck(end)
//		fmt.Println(reason)
//	}
//}
