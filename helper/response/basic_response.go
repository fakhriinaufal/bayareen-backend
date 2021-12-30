package response

type BasicResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}