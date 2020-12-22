package upnp

import (
	"io/ioutil"
	"log"

	// "log"
	"net/http"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type GetListOfPortMappings struct {
	upnp *Upnp
}

func (this *GetListOfPortMappings) Send() bool {
	request := this.buildRequest()
	response, _ := http.DefaultClient.Do(request)
	spew.Dump(response)
	resultBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {
		log.Println(string(resultBody))
		this.resolve(string(resultBody))
		return true
	}
	return false
}
func (this *GetListOfPortMappings) buildRequest() *http.Request {
	//请求头
	header := http.Header{}
	header.Set("Accept", "text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2")
	header.Set("SOAPAction", `"urn:schemas-upnp-org:service:WANIPConnection:2#GetListOfPortMappings"`)
	header.Set("Content-Type", "text/xml")
	header.Set("Connection", "Close")
	header.Set("Content-Length", "")

	//请求体
	body := Node{Name: "SOAP-ENV:Envelope",
		Attr: map[string]string{"xmlns:SOAP-ENV": `"http://schemas.xmlsoap.org/soap/envelope/"`,
			"SOAP-ENV:encodingStyle": `"http://schemas.xmlsoap.org/soap/encoding/"`}}
	childOne := Node{Name: `SOAP-ENV:Body`}
	childTwo := Node{Name: `m:GetListOfPortMappings`,
		Attr: map[string]string{"xmlns:m": `"urn:schemas-upnp-org:service:WANIPConnection:2"`}}
	childList1 := Node{Name: "NewStartPort", Content: "1"}
	childList2 := Node{Name: "NewEndPort", Content: "65535"}
	childList3 := Node{Name: "NewProtocol", Content: "UDP"}
	childList4 := Node{Name: "NewManage", Content: "1"}
	childList5 := Node{Name: "NewNumberOfPorts", Content: "65535"}

	childTwo.AddChild(childList1)
	childTwo.AddChild(childList2)
	childTwo.AddChild(childList3)
	childTwo.AddChild(childList4)
	childTwo.AddChild(childList5)
	childOne.AddChild(childTwo)
	body.AddChild(childOne)
	bodyStr := body.BuildXML()

	//请求
	request, _ := http.NewRequest("POST", "http://"+this.upnp.Gateway.Host+this.upnp.CtrlUrl,
		strings.NewReader(bodyStr))
	request.Header = header
	request.Header.Set("Content-Length", strconv.Itoa(len([]byte(bodyStr))))
	return request
}

func (this *GetListOfPortMappings) resolve(resultStr string) {
}