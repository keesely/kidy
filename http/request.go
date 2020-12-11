// request.go kee > 2020/12/01

package http

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	h "net/http"
	"okauth/utils"
	"strings"
)

type Headers struct {
	h.Header
}

type Request struct {
	request    *h.Request
	Host       string
	Path       string
	RequestURI string
	Proto      string
	Method     string
	Headers    Headers
	Values     utils.Values
	Files      Files
	Raw        []byte
	Query      utils.Values
	PostValues utils.Values
}

const (
	GET              = "GET"
	HEAD             = "HEAD"
	POST             = "POST"
	PUT              = "PUT"
	PATCH            = "PATCH" // RFC 5789
	DELETE           = "DELETE"
	CONNECT          = "CONNECT"
	OPTIONS          = "OPTIONS"
	TRACE            = "TRACE"
	defaultMaxMemory = 32 << 20 // 32 MB
)

func NewRequest(r *h.Request) *Request {
	headers := Headers{r.Header}

	request := &Request{
		request:    r,
		Host:       r.Host,
		Method:     r.Method,
		Proto:      r.Proto,
		Path:       r.URL.Path,
		RequestURI: r.URL.RequestURI(),
		Headers:    headers,
		Values:     utils.Values{},
	}
	request.setValues()
	return request
}

func (s *Request) setValues() {
	ct := s.Headers.Get("Content-Type")
	switch strings.Split(ct, ";")[0] {
	case "application/json":
		s.ParseJSON(&s.Values)
	case "application/x-www-form-urlencoded", "multipart/form-data":
		s.request.ParseMultipartForm(defaultMaxMemory)
		if strings.Split(ct, ";")[0] == "multipart/form-data" {
			s.Files = ParseFormFile(s.request)
		}
		if e := s.request.ParseForm(); e == nil {
			var formStr []string
			for k, v := range s.request.Form {
				formStr = append(formStr, fmt.Sprintf("%s=%s", k, v))
			}
			form := utils.Values{}
			ParseStr(strings.Join(formStr, "&"), form)
			for k, v := range form {
				s.Values.Set(k, v)
			}
		}
	}
	query := make(map[string]interface{})
	ParseStr(s.request.URL.RawQuery, query)
	for k, v := range query {
		if nil == s.Values.Get(k) {
			s.Values.Set(k, v)
		}
	}
}

func (s *Request) File(name string) (file FileStream, ok bool) {
	var e error
	if file, e = s.Files.Get(name); e == nil {
		ok = true
		return
	}
	return
}

func (s *Request) ParseJSON(v interface{}) error {
	return json.Unmarshal(s.GetRaw(), v)
}

func (s *Request) GetReq() *h.Request {
	return s.request
}

func (s *Request) GetRaw() []byte {
	if nil != s.Raw {
		return s.Raw
	}

	if con, e := ioutil.ReadAll(s.GetBody()); e == nil {
		s.Raw = con
	}
	return s.Raw
}

func (s *Request) GetQueryRaw() string {
	return s.request.URL.RawQuery
}

func (s *Request) GetBody() io.ReadCloser {
	return s.request.Body
}

func (s *Request) GetHeader(name string) string {
	return s.Headers.Get(name)
}

func (s *Request) GetValue(name string, def ...interface{}) interface{} {
	val, ok := s.Values[name]
	if !ok {
		return def
	}
	return val
}
