package ngssp

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

type SOAPType int

const (
	SOAP_DEF = iota + 1
	SOAP_SEL
)

const (
	defaultWaitTime       = 3 * time.Second
	defaultMaxWaitTime    = 10 * time.Second
	defaultDialTimeOut    = 3 * time.Second
	defaultRequestTimeOut = 3 * time.Second
)

type NgsspClient struct {
	cl      *http.Client
	baseUrl string
}

type option struct {
	dialTimeout time.Duration
	reqTimeout  time.Duration
	httpClient  *http.Client
}

type ClientOption func(*option)

func WithDialTimeout(d time.Duration) ClientOption {
	return func(o *option) {
		o.dialTimeout = d
	}
}

func WithRequestTimeout(d time.Duration) ClientOption {
	return func(o *option) {
		o.reqTimeout = d
	}
}

func WithHttpClient(cl *http.Client) ClientOption {
	return func(o *option) {
		o.httpClient = cl
	}
}

func NewClient(baseUrl string, options ...ClientOption) *NgsspClient {
	opt := option{
		dialTimeout: defaultDialTimeOut,
		reqTimeout:  defaultRequestTimeOut,
	}

	for _, op := range options {
		op(&opt)
	}

	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DialContext = (&net.Dialer{Timeout: opt.dialTimeout}).DialContext
	tr.TLSHandshakeTimeout = opt.dialTimeout

	var cl *http.Client
	if opt.httpClient != nil {
		cl = opt.httpClient
	} else {
		cl = &http.Client{
			Transport: tr,
			Timeout:   opt.reqTimeout,
		}
	}

	var trimUrlEnd func(string) string
	trimUrlEnd = func(str string) string {
		str = strings.TrimSuffix(str, "/")

		if strings.HasSuffix(str, "/") {
			return trimUrlEnd(str)
		}
		return str
	}

	baseUrl = trimUrlEnd(baseUrl)

	return &NgsspClient{
		cl:      cl,
		baseUrl: baseUrl,
	}
}

func (c *NgsspClient) GetPackageDef(pvr string) (*NgsspPackage, error) {
	res, err := c.callSoap(SOAP_DEF, pvr)
	if err != nil {
		return nil, err
	}

	return NewNgsspPackageDefDecoder(bytes.NewReader(res)).Decode()
}

func (c *NgsspClient) GetPackageSel(keyword string) (*NgsspVariantPackageSelector, error) {
	res, err := c.callSoap(SOAP_SEL, keyword)
	if err != nil {
		return nil, err
	}

	return NewNgsspPackageSelectorDecoder(bytes.NewReader(res)).Decode()
}

func (c *NgsspClient) callSoap(soapType SOAPType, key string) ([]byte, error) {
	var urlPath string
	var bodyStr string

	switch soapType {
	case SOAP_DEF:
		urlPath = c.baseUrl + packageDefinitionPath
		bodyStr = fmt.Sprintf(packageDefinitionRequestBody, key)
	case SOAP_SEL:
		urlPath = c.baseUrl + packageSelectorPath
		bodyStr = fmt.Sprintf(packageSelectorRequestBody, key)
	}

	req, err := http.NewRequest(http.MethodPost, urlPath, strings.NewReader(bodyStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	res, err := c.cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

var packageDefinitionPath = "/NgsspPackageDefinition/ProxyServices/PackageDefinitionPS"
var packageDefinitionRequestBody = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:pac="http://indosatooredoo.com/ngssp/schemas/packagedefinition">
    <soapenv:Header/>
    <soapenv:Body>
        <pac:NgsspPackageDefinitionRequest>
            <pac:keyName/>
            <pac:keyValue>%s</pac:keyValue>
        </pac:NgsspPackageDefinitionRequest>
    </soapenv:Body>
</soapenv:Envelope>
`

var packageSelectorPath = "/NgsspPackageSelector/ProxyServices/PackageSelectorPS"
var packageSelectorRequestBody = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:pac="http://indosatooredoo.com/ngssp/schemas/packageselector">
   <soapenv:Header/>
   <soapenv:Body>
      <pac:NgsspPackageSelectorRequest>
         <pac:keyName/>
         <pac:keyValue>%s</pac:keyValue>
      </pac:NgsspPackageSelectorRequest>
   </soapenv:Body>
</soapenv:Envelope>
`
