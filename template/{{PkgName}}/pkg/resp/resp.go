package resp

type Response struct {
	APIResult		APIResult `json:"api_result"`
	ServerRequest	ServerRequest `json:"server_request"`
	SystemInfo  	SystemInformation `json:"system_info"`
}
