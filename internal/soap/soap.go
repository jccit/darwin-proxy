package soap

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

const headerTemplate string = `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:typ="http://thalesgroup.com/RTTI/2013-11-28/Token/types" xmlns:ldb="http://thalesgroup.com/RTTI/2017-10-01/ldb/">
<soap:Header>
<typ:AccessToken>
<typ:TokenValue>{{.AccessKey}}</typ:TokenValue>
</typ:AccessToken>
</soap:Header>`

const darwinEndpoint = "https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb11.asmx"

type Header struct {
	AccessKey string
}

func GetSOAPHeader() string {
	t := template.New("req header")
	t, err := t.Parse(headerTemplate)

	if err != nil {
		return err.Error()
	}

	apiKey := os.Getenv("DARWIN_API_KEY")
	authData := Header{AccessKey: apiKey}

	var out bytes.Buffer
	t.Execute(&out, authData)

	return out.String()
}

func SendSOAPRequest(url string, soapStr string, r *http.Request) []byte {
	soapBody := []byte(soapStr)

	resp, err := http.Post(url, "text/xml", bytes.NewBuffer(soapBody))

	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func SendDarwinRequest(soapStr string, r *http.Request) []byte {
	return SendSOAPRequest(darwinEndpoint, GetSOAPHeader()+soapStr, r)
}
