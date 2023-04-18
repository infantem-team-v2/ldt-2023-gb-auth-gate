package dependency

import (
	"bank_api/config"
	"bank_api/pkg/thttp"
	"bank_api/pkg/tlogger"
	"github.com/sarulabs/di"
)

var dependencyMap = map[string]func(ctn di.Container) (interface{}, error){
	"config":     config.BuildConfig,
	"httpClient": thttp.BuildHttpClient,
	"logger":     tlogger.BuildLogger,
}

const TagDI = "di"
