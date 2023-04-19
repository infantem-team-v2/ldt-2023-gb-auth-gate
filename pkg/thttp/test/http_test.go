package test

import (
	"bank_api/config"
	"bank_api/pkg/thttp"
	"bank_api/pkg/tlogger"
	"github.com/sarulabs/di"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type ThttpClient struct {
	Config           *config.Config
	Logger           tlogger.TLogger
	httpClient       *http.Client
	responseRecorder *httptest.ResponseRecorder
}

//type mockTransport struct {
//	rt http.RoundTripper
//	fn func(*http.Request) (*http.Response, error)
//}

//func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
//	if m.fn != nil {
//		return m.fn(req)
//	}
//	return m.rt.RoundTrip(req)
//}

func TestQueryString(t *testing.T) {
	hc := &thttp.ThttpClient{}
	queryParams := map[string]interface{}{
		"user": "john",
		"id":   "123", //может быть только типа string
	}
	expectedQuery := "?user=john&id=123"
	query := hc.MakeQueryString(queryParams)
	if query != expectedQuery {
		t.Errorf("makeQuery(): expected '%s', got '%s'", expectedQuery, query)
	}
}

func TestRequest(t *testing.T) {
	builder, err := di.NewBuilder()
	if err != nil {
		t.Fatal(err)
	}

	cfg := config.Config{}
	logger := tlogger.TLogger{}
	recorder := httptest.NewRecorder()

	err = builder.Add(di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &cfg, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = builder.Add(di.Def{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return logger, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = builder.Add(di.Def{
		Name:  "httpClient",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &http.Client{}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = builder.Add(di.Def{
		Name:  "responseRecorder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return recorder, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = builder.Add(di.Def{
		Name:  "httpClientWrapper",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return ThttpClient{
				Config:           ctn.Get("config").(*config.Config),
				Logger:           ctn.Get("logger").(tlogger.TLogger),
				httpClient:       ctn.Get("httpClient").(*http.Client),
				responseRecorder: ctn.Get("responseRecorder").(*httptest.ResponseRecorder),
			}, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	container := builder.Build()

	expected := map[string]interface{}{
		"key": "value",
	}
	dest := map[string]interface{}{}
	reqParams := map[string]interface{}{
		"param": "value",
	}
	queryParams := map[string]interface{}{
		"queryParam": "value",
	}
	//body, _ := json.Marshal(reqParams)
	//resp := http.Response{
	//	StatusCode: 200,
	//	Body:       ioutil.NopCloser(bytes.NewReader(body)),
	//}
	//client := &http.Client{
	//	Transport: &mockTransport{
	//		fn: func(req *http.Request) (*http.Response, error) {
	//			return &resp, nil
	//		},
	//	},
	//}

	hc := container.Get("httpClientWrapper").(thttp.ThttpClient)

	result, statusCode, err := hc.Request("GET", "http://example.com", nil, reqParams, &dest, queryParams)
	if err != nil {
		t.Fatal(err)
	}
	if statusCode != 200 {
		t.Errorf("Request returned statusCode %d, expected 200", statusCode)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Request returned %v, expected %v", result, expected)
	}
}
