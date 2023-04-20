package test_http

import (
	"bank_api/internal/pkg/dependency"
	"bank_api/pkg/thttp"
	thttpHeaders "bank_api/pkg/thttp/headers"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// =======================HTTP=======================//
func TestRequest(t *testing.T) {
	tdc := dependency.NewDependencyContainer().BuildDependencies().BuildContainer()

	t.Run("non query params", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"key": "value",
		}
		dest := map[string]interface{}{}
		hc := tdc.ContainerInstance().Get("httpClient").(*thttp.ThttpClient)
		result, statusCode, err := hc.Request("GET", "https://api.twinklepick.com/api/transaction/get", nil, nil, &dest, nil)
		if err != nil {
			t.Fatal(err)
		}
		if statusCode != 200 {
			t.Errorf("Request returned statusCode %d, expected 200", statusCode)
		}
		assert.Equal(t, expectedResult, result)
	})
}

//api_key = 'jm5r4eutgpk2hqgt5dkrqn3ihf75wtvsfdzou4wf7feh7unmmsr5sv53yazfxuv8'
//api_secret = 'op9io5eyfne37xb2ad2ovyohb7oaqofjfp7yabuusz9tfbmm4z6hdkn7bx4gojfbq2fnuvxgecii5cocsrsu6x4ed9enui7z6t66hbfpdc4xcxmdca2xcb8t6tz2a8gz'

func TestLoad(b *testing.T) {
	tdc := dependency.NewDependencyContainer().BuildDependencies().BuildContainer()
	httpClient := tdc.ContainerInstance().Get("httpClient").(*thttp.ThttpClient)
	body := map[string]string{"tracker_id": "8599ea25812430e06f19ffaa5969816c046153387e311f867fa5e1f44f651dad58f58a7effbbcb0200b6b7a38f2137fbe5ef50c22db65ea13090ff4261b15677"}
	headers, err := thttpHeaders.MakeAuthHeaders(body,
		"jm5r4eutgpk2hqgt5dkrqn3ihf75wtvsfdzou4wf7feh7unmmsr5sv53yazfxuv8",
		"op9io5eyfne37xb2ad2ovyohb7oaqofjfp7yabuusz9tfbmm4z6hdkn7bx4gojfbq2fnuvxgecii5cocsrsu6x4ed9enui7z6t66hbfpdc4xcxmdca2xcb8t6tz2a8gz",
		thttp.POST,
	)
	if err != nil {

		b.Fatal(err)
	}
	timeSlice := make([]int64, 9, 10000)
	var total int64 = 0
	c := make(chan int64)
	for i := 0; i < 10000; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				start := time.Now().UnixNano()
				if r, s, e := httpClient.Request(thttp.POST, "https://test.onecrypto.pro/api/transaction/get", headers, body, nil, nil); e != nil {
					b.Log(r, s)
				}
				end := time.Now().UnixNano()
				timeSlice = append(timeSlice, end-start)
				c <- end - start
			}

		}()

	}
	total += <-c
	b.Logf("Avg time: %d", total/int64(len(timeSlice)))
}

func TestMakeQueryString(t *testing.T) {
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
