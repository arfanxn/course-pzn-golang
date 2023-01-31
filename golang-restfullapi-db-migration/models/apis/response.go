package apis

type Response struct {
	Code   int16  `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}
