package common

type Params struct {
	UserId uint64 `json:"-"`
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
