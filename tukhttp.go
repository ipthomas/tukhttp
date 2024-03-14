package tukhttp

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ipthomas/tukcnst"
	"github.com/ipthomas/tukutil"
	"github.com/jrivets/log4g"
)

var stl = log4g.GetLogger("2T-UIS")

type HTTPRequest struct {
	Act          string
	Version      int
	Server       string
	Header       http.Header
	Method       string
	URL          string
	X_Api_Key    string
	X_Api_Secret string
	PID_OID      string
	PID          string
	SOAPAction   string
	ContentType  string
	Body         []byte
	Timeout      int
	StatusCode   int
	Response     []byte
	DebugMode    bool
}
type TukHTTPInterface interface {
	newRequest() error
}

func NewRequest(i TukHTTPInterface) error {
	return i.newRequest()
}
func (i *HTTPRequest) newRequest() error {
	var err error
	if i.Timeout == 0 {
		i.Timeout = 15
	}
	if err = i.sendHttpRequest(); err != nil {
		stl.Error(err.Error())
	}
	i.logRsp(i.StatusCode, string(i.Response))
	return err

}
func (i *HTTPRequest) sendHttpRequest() error {
	var err error
	var req *http.Request = new(http.Request)
	var rsp *http.Response
	var bytes []byte
	req.Header.Add(tukcnst.ACCEPT, tukcnst.ALL)
	req.Header.Add(tukcnst.CONTENT_TYPE, i.ContentType)
	req.Header.Add(tukcnst.CONNECTION, tukcnst.KEEP_ALIVE)
	if i.X_Api_Key != "" && i.X_Api_Secret != "" {
		req.Header.Add("X-API-KEY", i.X_Api_Key)
		req.Header.Add("X-API-SECRET", i.X_Api_Secret)
	}
	if i.SOAPAction != "" {
		req.Header.Set(tukcnst.SOAP_ACTION, i.SOAPAction)
	}
	i.logReq(req.Header, i.URL, string(i.Body))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.Timeout)*time.Second)
	defer cancel()
	if i.Method == http.MethodPost && string(i.Body) != "" {
		req, err = http.NewRequest(i.Method, i.URL, strings.NewReader(string(i.Body)))
	} else {
		req, err = http.NewRequest(i.Method, i.URL, nil)
	}
	if err == nil {
		if rsp, err = http.DefaultClient.Do(req.WithContext(ctx)); err == nil {
			defer rsp.Body.Close()
			if bytes, err = io.ReadAll(rsp.Body); err == nil {
				i.StatusCode = rsp.StatusCode
				i.Response = bytes
			}
		}
	}
	return err
}
func (i *HTTPRequest) logReq(headers interface{}, url string, body string) {
	if i.DebugMode {
		stl.Debug("HTTP Request Headers")
		tukutil.Log(headers)
		stl.Debug("\nHTTP Request\n-- URL = %s", url)
		if body != "" {
			stl.Debug("\n-- Body:\n%s", body)
		}
	}
}
func (i *HTTPRequest) logRsp(statusCode int, response string) {
	if i.DebugMode {
		stl.Debug("HTML Response - Status Code = %v\n-- Response--\n%s", statusCode, response)
	}
}
