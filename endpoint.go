package cloudrequests

import(
	"fmt"
)

type Endpoint struct{
	Url string `json:"url"`
	Name string `json:"name"`
        Header string `json:"header"`
        Data string `json:"data"`
//        data DataType `json:"data"`
        Method string `json:"method"`
	Proxy string `json:"proxy"`
	Timeout int `json:"timeout"`
	Expected Response `json:"expected"`
}

//func (e Endpoint) Data() string {
//	return e.data.plain
//}

func (dt DataType) UnmarshalJSON (buf []byte) error {
	fmt.Println("unmarshal: "+string(buf))	
	dt.plain = string(buf)

	return nil
}

type DataType struct{
	plain string	
}

//func NewEndpoint(url string) Endpoint{
//	var resp Response
//	return Endpoint{url,"","","","Get","",600,resp}
//}

type Response struct{
	Status int `json:"status"`
	Header string `json:string`
	Data string `json:"data"`
	ErrorMessage string
}
