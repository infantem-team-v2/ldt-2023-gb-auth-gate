package test

import (
	"bank_api/config"
	"bank_api/pkg/thttp"
	"bank_api/pkg/tlogger"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//=======================HTTP=======================//

func TestRequest(t *testing.T) {
	builder, err := di.NewBuilder()
	if err != nil {
		t.Fatal(err)
	}

	err = builder.Add(di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return config.NewConfig(), nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = builder.Add(di.Def{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return tlogger.NewLogger(), nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Fatal(err)
	}
	err = builder.Add(di.Def{
		Name:  "httpClientWrapper",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return thttp.ThttpClient{
				Config: ctn.Get("config").(*config.Config),
				Logger: ctn.Get("logger").(tlogger.Logger),
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
	hc := container.Get("httpClientWrapper").(thttp.ThttpClient)

	result, statusCode, err := hc.Request("GET", "https://api.worldchange.cc/api/currencies/get_from", nil, nil, &dest, nil)
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

func TestThttpClient_MakeQueryString(t *testing.T) {
	hc := thttp.ThttpClient{}

	t.Run("non-empty query params", func(t *testing.T) {
		queryParams := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}

		expectedResult := "?key1=value1&key2=value2"

		result := hc.MakeQueryString(queryParams)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("empty query params", func(t *testing.T) {
		queryParams := map[string]interface{}{}

		expectedResult := ""

		result := hc.MakeQueryString(queryParams)

		assert.Equal(t, expectedResult, result)
	})
}
